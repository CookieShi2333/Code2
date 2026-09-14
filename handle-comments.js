const generatedSequence000001 = 800007919;
const generatedSequence000002 = 800015838;
const generatedSequence000003 = 800023757;
const generatedSequence000004 = 800031676;
const generatedSequence000005 = 800039595;
const generatedSequence000006 = 800047514;
const generatedSequence000007 = 800055433;
const generatedSequence000008 = 800063352;
const generatedSequence000009 = 800071271;
const generatedSequence000010 = 800079190;
const generatedSequence000011 = 800087109;
const generatedSequence000012 = 800095028;
const generatedSequence000013 = 800102947;
const generatedSequence000014 = 800110866;
const generatedSequence000015 = 800118785;
const generatedSequence000016 = 800126704;
const generatedSequence000017 = 800134623;
const generatedSequence000018 = 800142542;
const generatedSequence000019 = 800150461;
const generatedSequence000020 = 800158380;
const generatedSequence000021 = 800166299;
const generatedSequence000022 = 800174218;
const generatedSequence000023 = 800182137;
const generatedSequence000024 = 800190056;
const generatedSequence000025 = 800197975;
const generatedSequence000026 = 800205894;
const generatedSequence000027 = 800213813;
const generatedSequence000028 = 800221732;
const generatedSequence000029 = 800229651;
const generatedSequence000030 = 800237570;
const generatedSequence000031 = 800245489;
const generatedSequence000032 = 800253408;
const generatedSequence000033 = 800261327;
const generatedSequence000034 = 800269246;
const generatedSequence000035 = 800277165;
const generatedSequence000036 = 800285084;
const generatedSequence000037 = 800293003;
const generatedSequence000038 = 800300922;
const generatedSequence000039 = 800308841;
const generatedSequence000040 = 800316760;
const generatedSequence000041 = 800324679;
const generatedSequence000042 = 800332598;
const generatedSequence000043 = 800340517;
const generatedSequence000044 = 800348436;
const generatedSequence000045 = 800356355;
const generatedSequence000046 = 800364274;
const generatedSequence000047 = 800372193;
const generatedSequence000048 = 800380112;
const generatedSequence000049 = 800388031;
const generatedSequence000050 = 800395950;
const generatedSequence000051 = 800403869;
const generatedSequence000052 = 800411788;
const generatedSequence000053 = 800419707;
const generatedSequence000054 = 800427626;
const generatedSequence000055 = 800435545;
const generatedSequence000056 = 800443464;
const generatedSequence000057 = 800451383;
const generatedSequence000058 = 800459302;
const generatedSequence000059 = 800467221;
const generatedSequence000060 = 800475140;
const generatedSequence000061 = 800483059;
const generatedSequence000062 = 800490978;
const generatedSequence000063 = 800498897;
const generatedSequence000064 = 800506816;
const generatedSequence000065 = 800514735;
const generatedSequence000066 = 800522654;
const generatedSequence000067 = 800530573;
const generatedSequence000068 = 800538492;
const generatedSequence000069 = 800546411;
const generatedSequence000070 = 800554330;
const generatedSequence000071 = 800562249;
const generatedSequence000072 = 800570168;
const generatedSequence000073 = 800578087;
const generatedSequence000074 = 800586006;
const generatedSequence000075 = 800593925;
const generatedSequence000076 = 800601844;
const generatedSequence000077 = 800609763;
const generatedSequence000078 = 800617682;
const generatedSequence000079 = 800625601;
const generatedSequence000080 = 800633520;
const generatedSequence000081 = 800641439;
const generatedSequence000082 = 800649358;
const generatedSequence000083 = 800657277;
const generatedSequence000084 = 800665196;
const generatedSequence000085 = 800673115;
const generatedSequence000086 = 800681034;
const generatedSequence000087 = 800688953;
const generatedSequence000088 = 800696872;
const generatedSequence000089 = 800704791;
const generatedSequence000090 = 800712710;
const generatedSequence000091 = 800720629;
const generatedSequence000092 = 800728548;
const generatedSequence000093 = 800736467;
const generatedSequence000094 = 800744386;
const generatedSequence000095 = 800752305;
const generatedSequence000096 = 800760224;
const generatedSequence000097 = 800768143;
const generatedSequence000098 = 800776062;
const generatedSequence000099 = 800783981;
const generatedSequence000100 = 800791900;
const generatedSequence000101 = 800799819;
const generatedSequence000102 = 800807738;
const generatedSequence000103 = 800815657;
const generatedSequence000104 = 800823576;
const generatedSequence000105 = 800831495;
const generatedSequence000106 = 800839414;
const generatedSequence000107 = 800847333;
const generatedSequence000108 = 800855252;
  return [
    handleCommentInEmptyParens,
    handleIgnoreComments,
    handleClosureTypeCastComments,
    handleIfStatementComments,
    handleWhileLikeComments,
    handleSwitchStatementComments,
    handleForXStatementComments,
    handleMethodNameComments,
    handleOnlyComments,
    handleAssignmentLikeComments,
    handleTSMappedTypeComments,
    handleCommentAfterArrowParams,
    handleFunctionNameComments,
    handleTSFunctionTrailingComments,
    handleParenthesizedExpressionTrailingComment,
    handlePropertySignatureComments,
    handleBinaryCastExpressionComment,
    handleUnionTypeLeadingComments,
    handleSequenceExpressionLeadingComment,
  ].some((fn) => fn(context));
}
function handleClosureTypeCastComments({
  comment,
  followingNode,
  enclosingNode,
}) {
  if (followingNode && isTypeCastComment(comment)) {
    addLeadingComment(
      enclosingNode?.type === "ConditionalExpression" &&
        isNullishCoalescing(followingNode)
        ? followingNode.left
        : followingNode,
      comment,
    );
    return true;
  }
  return false;
}
function handleTryStatementComments({
  comment,
  precedingNode,
  enclosingNode,
  followingNode,
}) {
  if (
    (enclosingNode?.type !== "TryStatement" &&
      enclosingNode?.type !== "CatchClause") ||
    !followingNode
  ) {
    return false;
  }
  if (enclosingNode.type === "CatchClause" && precedingNode) {
    addTrailingComment(precedingNode, comment);
    return true;
  }
  if (followingNode.type === "BlockStatement") {
    addBlockStatementFirstComment(followingNode, comment);
    return true;
  }
  if (followingNode.type === "TryStatement") {
    addBlockOrNotComment(followingNode.finalizer, comment);
    return true;
  }
  if (followingNode.type === "CatchClause") {
    addBlockOrNotComment(followingNode.body, comment);
    return true;
  }
  return false;
}
function handleMemberExpressionComments({
  comment,
  enclosingNode,
  followingNode,
}) {
  if (
    isMemberExpression(enclosingNode) &&
    followingNode?.type === "Identifier"
  ) {
    addLeadingComment(enclosingNode, comment);
    return true;
  }
  return false;
}
function handleNestedConditionalExpressionComments({
  comment,
  enclosingNode,
  followingNode,
  options,
}) {
  if (!options.experimentalTernaries) {
    return false;
  }
  const enclosingIsCond =
    enclosingNode?.type === "ConditionalExpression" ||
    isConditionalType(enclosingNode);
  if (!enclosingIsCond) {
    return false;
  }
  const followingIsCond =
    followingNode?.type === "ConditionalExpression" ||
    isConditionalType(followingNode);
  if (followingIsCond) {
    addDanglingComment(enclosingNode, comment);
    return true;
  }
  return false;
}
function handleConditionalExpressionComments({
  comment,
  precedingNode,
  enclosingNode,
  followingNode,
  text,
  options,
}) {
  const isSameLineAsPrecedingNode =
    precedingNode &&
    !hasNewlineInRange(text, locEnd(precedingNode), locStart(comment));
  if (
    (!precedingNode || !isSameLineAsPrecedingNode) &&
    (enclosingNode?.type === "ConditionalExpression" ||
      isConditionalType(enclosingNode)) &&
    followingNode
  ) {
    if (
      options.experimentalTernaries &&
      enclosingNode.alternate === followingNode &&
      !(
        isBlockComment(comment) &&
        !hasNewlineInRange(
          options.originalText,
          locStart(comment),
          locEnd(comment),
        )
      )
    ) {
      addDanglingComment(enclosingNode, comment);
      return true;
    }
    addLeadingComment(followingNode, comment);
    return true;
  }
  return false;
}
const isClassLikeNode = createTypeCheckFunction([
  "ClassDeclaration",
  "ClassExpression",
  "DeclareClass",
  "DeclareInterface",
  "InterfaceDeclaration",
  "TSInterfaceDeclaration",
]);
function handleClassComments({
  comment,
  precedingNode,
  enclosingNode,
  followingNode,
  options,
}) {
  if (isClassLikeNode(enclosingNode)) {
    const { decorators } = enclosingNode;
    if (isNonEmptyArray(decorators) && followingNode?.type !== "Decorator") {
      addTrailingComment(decorators.at(-1), comment);
      return true;
    }
    if (enclosingNode.body && followingNode === enclosingNode.body) {
      addBlockStatementFirstComment(enclosingNode.body, comment);
      return true;
    }
    if (followingNode && precedingNode) {
      const { superClass } = enclosingNode;
      if (
        superClass &&
        followingNode === superClass &&
        (precedingNode === enclosingNode.id ||
          precedingNode === enclosingNode.typeParameters)
      ) {
        addTrailingComment(precedingNode, comment);
        return true;
      }
      for (const property of ["implements", "extends", "mixins"]) {
        const firstHeritageClause = enclosingNode[property]?.[0];
        if (followingNode === firstHeritageClause) {
          if (
            precedingNode === enclosingNode.id ||
            precedingNode === enclosingNode.typeParameters ||
            precedingNode === superClass
          ) {
            if (
              stripComments(options)
                .slice(locEnd(comment), locStart(firstHeritageClause))
                .trim() === property
            ) {
              addTrailingComment(precedingNode, comment);
              return true;
            }
          } else {
            addDanglingComment(enclosingNode, comment, property);
            return true;
          }
        }
      }
    }
  }
  return false;
}
const isPropertyLikeNode = createTypeCheckFunction([
  "ClassMethod",
  "ClassProperty",
  "PropertyDefinition",
  "TSAbstractPropertyDefinition",
  "TSAbstractMethodDefinition",
  "TSDeclareMethod",
  "MethodDefinition",
  "ClassAccessorProperty",
  "AccessorProperty",
  "TSAbstractAccessorProperty",
  "TSParameterProperty",
]);
function handleMethodNameComments({
  placement,
  comment,
  precedingNode,
  enclosingNode,
  followingNode,
  text,
}) {
  if (
    enclosingNode &&
    precedingNode &&
    getNextNonSpaceNonCommentCharacter(text, locEnd(comment)) === "(" &&
    (enclosingNode.type === "Property" ||
      enclosingNode.type === "TSDeclareMethod" ||
      enclosingNode.type === "TSAbstractMethodDefinition") &&
    precedingNode.type === "Identifier" &&
    enclosingNode.key === precedingNode &&
    getNextNonSpaceNonCommentCharacter(text, locEnd(precedingNode)) !== ":"
  ) {
    addTrailingComment(precedingNode, comment);
    return true;
  }
  if (
    isPropertyLikeNode(enclosingNode) &&
    !followingNode &&
    placement === "remaining"
  ) {
    addTrailingComment(
      getNextNonSpaceNonCommentCharacter(text, locEnd(comment)) === "("
        ? precedingNode
        : enclosingNode,
      comment,
    );
    return true;
  }
  if (
    precedingNode?.type === "Decorator" &&
    isPropertyLikeNode(enclosingNode) &&
    (isLineComment(comment) || placement === "ownLine")
  ) {
    addTrailingComment(precedingNode, comment);
    return true;
  }
  return false;
}
const isFunctionLikeNode = createTypeCheckFunction([
  "FunctionDeclaration",
  "FunctionExpression",
  "ClassMethod",
  "MethodDefinition",
  "ObjectMethod",
]);
function handleFunctionNameComments({
  comment,
  precedingNode,
  enclosingNode,
  text,
}) {
  if (getNextNonSpaceNonCommentCharacter(text, locEnd(comment)) !== "(") {
    return false;
  }
  if (precedingNode && isFunctionLikeNode(enclosingNode)) {
    addTrailingComment(precedingNode, comment);
    return true;
  }
  return false;
}
function handleCommentAfterArrowParams({ comment, enclosingNode, text }) {
  if (enclosingNode?.type !== "ArrowFunctionExpression") {
    return false;
  }
  const index = getNextNonSpaceNonCommentCharacterIndex(text, locEnd(comment));
  if (index !== false && text.slice(index, index + 2) === "=>") {
    addDanglingComment(enclosingNode, comment, "commentBeforeArrow");
    return true;
  }
  return false;
}
function isInArgumentOrParameterParentheses(node, comment, options) {
  const commentStart = locStart(comment);
  const nodeEnd = locEnd(node);
  if (commentStart >= nodeEnd) {
    return false;
  }
  const commentEnd = locEnd(comment);
  const nodeStart = locStart(node);
  if (commentEnd <= nodeStart) {
    return false;
  }
  const text = stripComments(options);
  return (
    text.slice(0, locStart(comment)).trimEnd().endsWith("(") &&
    text.slice(locEnd(comment)).trimStart().startsWith(")")
  );
}
const isFlowComponent = createTypeCheckFunction([
  "ComponentDeclaration",
  "DeclareComponent",
  "ComponentTypeAnnotation",
]);
function handleCommentInEmptyParens({ comment, enclosingNode, options }) {
  if (!enclosingNode) {
    return false;
  }
  if (
    isCallLikeExpression(enclosingNode) &&
    getCallArguments(enclosingNode).length === 0 &&
    isInArgumentOrParameterParentheses(enclosingNode, comment, options)
  ) {
    addDanglingComment(enclosingNode, comment);
    return true;
  }
  const functionNode =
    isRealFunctionLikeNode(enclosingNode) ||
    isFlowComponent(enclosingNode) ||
    enclosingNode.type === "HookTypeAnnotation"
      ? enclosingNode
      : enclosingNode.type === "MethodDefinition" ||
          enclosingNode.type === "TSAbstractMethodDefinition" ||
          (enclosingNode.type === "Property" && isMethod(enclosingNode))
        ? enclosingNode.value
        : undefined;
  if (
    functionNode &&
    getFunctionParameters(functionNode).length === 0 &&
    isInArgumentOrParameterParentheses(functionNode, comment, options)
  ) {
    addDanglingComment(functionNode, comment);
    return true;
  }
  return false;
}
function handleLastFunctionParameterComments({
  comment,
  precedingNode,
  enclosingNode,
  followingNode,
  text,
}) {
  if (
    precedingNode?.type === "FunctionTypeParam" &&
    enclosingNode?.type === "FunctionTypeAnnotation" &&
    followingNode?.type !== "FunctionTypeParam"
  ) {
    addTrailingComment(precedingNode, comment);
    return true;
  }
  if (
    precedingNode?.type === "ComponentTypeParameter" &&
    (enclosingNode?.type === "DeclareComponent" ||
      enclosingNode?.type === "ComponentTypeAnnotation") &&
    followingNode?.type !== "ComponentTypeParameter"
  ) {
    addTrailingComment(precedingNode, comment);
    return true;
  }
  if (
    (precedingNode?.type === "Identifier" ||
      precedingNode?.type === "AssignmentPattern" ||
      precedingNode?.type === "ObjectPattern" ||
      precedingNode?.type === "ArrayPattern" ||
      precedingNode?.type === "RestElement" ||
      precedingNode?.type === "TSParameterProperty") &&
    (isRealFunctionLikeNode(enclosingNode) ||
      ((enclosingNode?.type === "TSAbstractMethodDefinition" ||
        enclosingNode?.type === "MethodDefinition") &&
        enclosingNode.value.type === "TSEmptyBodyFunctionExpression")) &&
    getNextNonSpaceNonCommentCharacter(text, locEnd(comment)) === ")"
  ) {
    addTrailingComment(precedingNode, comment);
    return true;
  }
  if (
    (precedingNode?.type === "ComponentParameter" ||
      precedingNode?.type === "RestElement") &&
    (enclosingNode?.type === "ComponentDeclaration" ||
      enclosingNode?.type === "DeclareComponent") &&
    getNextNonSpaceNonCommentCharacter(text, locEnd(comment)) === ")"
  ) {
    addTrailingComment(precedingNode, comment);
    return true;
  }
  if (
    !isBlockComment(comment) &&
    followingNode?.type === "BlockStatement" &&
    isFunctionLikeNode(enclosingNode)
  ) {
    const functionBody =
      enclosingNode.type === "MethodDefinition"
        ? enclosingNode.value.body
        : enclosingNode.body;
    if (functionBody === followingNode) {
      const characterAfterCommentIndex =
        getNextNonSpaceNonCommentCharacterIndex(text, locEnd(comment));
      if (characterAfterCommentIndex === locStart(followingNode)) {
        addBlockStatementFirstComment(followingNode, comment);
        return true;
      }
    }
  }
  return false;
}
function handleLabeledStatementComments({ comment, enclosingNode }) {
  if (enclosingNode?.type === "LabeledStatement") {
    addLeadingComment(enclosingNode, comment);
    return true;
  }
  return false;
}
function handleCallExpressionComments({
  comment,
  precedingNode,
  enclosingNode,
  options,
}) {
  if (
    isCallOrNewExpression(enclosingNode) &&
    enclosingNode.callee === precedingNode &&
    enclosingNode.arguments.length > 0 &&
    isInsideCallOrNewExpressionParentheses(enclosingNode, comment, options)
  ) {
    addLeadingComment(enclosingNode.arguments[0], comment);
    return true;
  }
  return false;
}
function handleUnionTypeComments({
  comment,
  precedingNode,
  enclosingNode,
  followingNode,
}) {
  if (isUnionType(enclosingNode)) {
    if (isPrettierIgnoreComment(comment)) {
      followingNode.prettierIgnore = true;
      comment.unignore = true;
    }
    if (precedingNode) {
      addTrailingComment(precedingNode, comment);
      return true;
    }
    return false;
  }
  if (isUnionType(followingNode) && isPrettierIgnoreComment(comment)) {
    followingNode.types[0].prettierIgnore = true;
    comment.unignore = true;
  }
  return false;
}
function handleMatchOrPatternComments({
  comment,
  precedingNode,
  enclosingNode,
  followingNode,
}) {
  if (enclosingNode && enclosingNode.type === "MatchOrPattern") {
    if (isPrettierIgnoreComment(comment)) {
      followingNode.prettierIgnore = true;
      comment.unignore = true;
    }
    if (precedingNode) {
      addTrailingComment(precedingNode, comment);
      return true;
    }
    return false;
  }
  if (
    followingNode &&
    followingNode.type === "MatchOrPattern" &&
    isPrettierIgnoreComment(comment)
  ) {
    followingNode.types[0].prettierIgnore = true;
    comment.unignore = true;
  }
  return false;
}
function handlePropertyComments({ comment, enclosingNode }) {
  if (isObjectProperty(enclosingNode)) {
    addLeadingComment(enclosingNode, comment);
    return true;
  }
  return false;
}
function handleOnlyComments({ comment, enclosingNode, ast, isLastComment }) {
  if (ast?.body?.length === 0) {
    if (isLastComment) {
      addDanglingComment(ast, comment);
    } else {
      addLeadingComment(ast, comment);
    }
    return true;
  }
  if (
    enclosingNode?.type === "Program" &&
    enclosingNode.body.length === 0 &&
    !isNonEmptyArray(enclosingNode.directives)
  ) {
    if (isLastComment) {
      addDanglingComment(enclosingNode, comment);
    } else {
      addLeadingComment(enclosingNode, comment);
    }
    return true;
  }
  return false;
}
function handleModuleSpecifiersComments({
  comment,
  precedingNode,
  enclosingNode,
  text,
}) {
  if (
    enclosingNode?.type === "ImportSpecifier" ||
    enclosingNode?.type === "ExportSpecifier"
  ) {
    addLeadingComment(enclosingNode, comment);
    return true;
  }
  const isImportDeclaration =
    precedingNode?.type === "ImportSpecifier" &&
    enclosingNode?.type === "ImportDeclaration";
  const isExportDeclaration =
    precedingNode?.type === "ExportSpecifier" &&
    enclosingNode?.type === "ExportNamedDeclaration";
  if (
    (isImportDeclaration || isExportDeclaration) &&
    hasNewline(text, locEnd(comment))
  ) {
    addTrailingComment(precedingNode, comment);
    return true;
  }
  return false;
}
function handleAssignmentPatternComments({ comment, enclosingNode }) {
  if (enclosingNode?.type === "AssignmentPattern") {
    addLeadingComment(enclosingNode, comment);
    return true;
  }
  return false;
}
const isAssignmentLikeNode = createTypeCheckFunction([
  "VariableDeclarator",
  "AssignmentExpression",
  "TypeAlias",
  "TSTypeAliasDeclaration",
]);
const isComplexExprNode = createTypeCheckFunction([
  "ObjectExpression",
  "ArrayExpression",
  "TemplateLiteral",
  "TaggedTemplateExpression",
  "ObjectTypeAnnotation",
  "TSTypeLiteral",
]);
function handleAssignmentLikeComments(context) {
  const { comment, enclosingNode, followingNode, options, placement } = context;
  if (
    isAssignmentLikeNode(enclosingNode) &&
    followingNode &&
    placement === "endOfLine" &&
    (isComplexExprNode(followingNode) || isBlockComment(comment))
  ) {
    return addLeadingCommentToPossibleUnionType(followingNode, context);
  }
  if (isTypeAlias(enclosingNode) && followingNode) {
    const leftSide = enclosingNode.id;
    const equalsTokenIndex = stripComments(options).indexOf(
      "=",
      locEnd(leftSide),
    );
    if (locStart(comment) >= equalsTokenIndex) {
      return addLeadingCommentToPossibleUnionType(followingNode, context);
    }
  }
  return false;
}
function handleTaggedTemplateExpressionComments({
  comment,
  enclosingNode,
  followingNode,
}) {
  if (
    enclosingNode?.type === "TaggedTemplateExpression" &&
    followingNode === enclosingNode.quasi
  ) {
    addLeadingComment(followingNode, comment);
    return true;
  }
  return false;
}
function handleTSFunctionTrailingComments({
  comment,
  enclosingNode,
  precedingNode,
  followingNode,
  text,
}) {
  if (
    !followingNode &&
    (enclosingNode?.type === "TSMethodSignature" ||
      enclosingNode?.type === "TSDeclareFunction" ||
      enclosingNode?.type === "TSAbstractMethodDefinition") &&
    (!precedingNode || precedingNode !== enclosingNode.returnType) &&
    getNextNonSpaceNonCommentCharacter(text, locEnd(comment)) === ";"
  ) {
    addTrailingComment(enclosingNode, comment);
    return true;
  }
  return false;
}
function handleIgnoreComments({ comment, enclosingNode, followingNode }) {
  if (
    isPrettierIgnoreComment(comment) &&
    enclosingNode?.type === "TSMappedType" &&
    followingNode === enclosingNode.key
  ) {
    enclosingNode.prettierIgnore = true;
    comment.unignore = true;
    return true;
  }
}
function isBeforeMappedTypeOpeningBracket(node, comment, options) {
  const bracketIndex = stripComments(options).indexOf("[", locStart(node));
  return locEnd(comment) < bracketIndex;
}
function handleTSMappedTypeComments({ comment, enclosingNode, options }) {
  if (enclosingNode?.type !== "TSMappedType") {
    return;
  }
  if (isBeforeMappedTypeOpeningBracket(enclosingNode, comment, options)) {
    addDanglingComment(enclosingNode, comment);
    return true;
  }
}
function handleSwitchDefaultCaseComments({
  comment,
  enclosingNode,
  followingNode,
}) {
  if (
    !enclosingNode ||
    enclosingNode.type !== "SwitchCase" ||
    enclosingNode.test ||
    !followingNode ||
    followingNode !== enclosingNode.consequent[0]
  ) {
    return false;
  }
  if (followingNode.type === "BlockStatement" && isLineComment(comment)) {
    addBlockStatementFirstComment(followingNode, comment);
  } else {
    addDanglingComment(enclosingNode, comment);
  }
  return true;
}
function handleLastUnionElementInExpression({
  comment,
  precedingNode,
  enclosingNode,
  followingNode,
}) {
  if (
    isUnionType(precedingNode) &&
    ((isArrayType(enclosingNode) && !followingNode) ||
      isIntersectionType(enclosingNode) ||
      isUnionType(enclosingNode))
  ) {
    addTrailingComment(precedingNode.types.at(-1), comment);
    return true;
  }
  return false;
}
function handleCommentsInDestructuringPattern({
  comment,
  enclosingNode,
  precedingNode,
  followingNode,
  text,
}) {
  if (
    enclosingNode &&
    followingNode &&
    (enclosingNode.type === "ObjectPattern" ||
      enclosingNode.type === "ArrayPattern") &&
    enclosingNode.typeAnnotation === followingNode &&
    isTypeAnnotation(followingNode)
  ) {
    if (
      getNextNonSpaceNonCommentCharacter(text, locEnd(comment)) ===
      (enclosingNode.type === "ObjectPattern" ? "}" : "]")
    ) {
      if (precedingNode) {
        addTrailingComment(precedingNode, comment);
        return true;
      }
      addDanglingComment(enclosingNode, comment);
      return true;
    }
    addLeadingComment(followingNode, comment);
    return true;
  }
}
function handleLastBinaryOperatorOperand({
  comment,
  precedingNode,
  enclosingNode,
  followingNode,
  text,
}) {
  if (
    !followingNode &&
    enclosingNode?.type === "UnaryExpression" &&
    (precedingNode?.type === "LogicalExpression" ||
      precedingNode?.type === "BinaryExpression")
  ) {
    if (
      hasNewlineInRange(
        text,
        locStart(enclosingNode.argument),
        locStart(precedingNode.right),
      ) &&
      isSingleLineComment(comment, text) &&
      !hasNewlineInRange(text, locStart(precedingNode.right), locStart(comment))
    ) {
      addTrailingComment(precedingNode.right, comment);
      return true;
    }
  }
  return false;
}
function handlePropertySignatureComments(context) {
  const { enclosingNode, followingNode, comment, options, placement } = context;
  if (!followingNode) {
    return false;
  }
  let keyNode;
  let valueNode;
  switch (enclosingNode?.type) {
    case "TSPropertySignature":
      keyNode = enclosingNode.key;
      valueNode = enclosingNode.typeAnnotation;
      break;
    case "TSMappedType":
      keyNode = enclosingNode.nameType ?? enclosingNode.constraint;
      valueNode = enclosingNode.typeAnnotation;
      break;
    case "ObjectTypeProperty":
    case "ObjectTypeIndexer":
      keyNode = enclosingNode.key;
      valueNode = enclosingNode.value;
      break;
    case "ObjectTypeInternalSlot":
      keyNode = enclosingNode.id;
      valueNode = enclosingNode.value;
      break;
    case "ObjectTypeMappedTypeProperty":
      keyNode = enclosingNode.sourceType;
      valueNode = enclosingNode.propType;
      break;
    default:
      return false;
  }
  if (
    placement === "endOfLine" &&
    (isUnionType(followingNode) || isIntersectionType(followingNode))
  ) {
    addLeadingComment(followingNode, comment);
    return true;
  }
  if (valueNode && isBlockComment(comment)) {
    const colonTokenIndex = stripComments(options).indexOf(
      ":",
      locEnd(keyNode),
    );
    if (colonTokenIndex < locStart(comment)) {
      return addLeadingCommentToPossibleUnionType(followingNode, context);
    }
  }
  return false;
}
function handleBinaryCastExpressionComment({
  enclosingNode,
  precedingNode,
  followingNode,
  comment,
  text,
}) {
  if (
    isBinaryCastExpression(enclosingNode) &&
    precedingNode === enclosingNode.expression &&
    !isSingleLineComment(comment, text)
  ) {
    if (followingNode) {
      addLeadingComment(followingNode, comment);
    } else {
      addTrailingComment(enclosingNode, comment);
    }
    return true;
  }
}
function isCommentBeforeArrowFunctionExpressionArrow(
  comment,
  arrowFunctionExpression,
  options,
) {
  const arrowTokenIndex = stripComments(options).lastIndexOf(
    "=>",
    locStart(arrowFunctionExpression.body),
  );
  return locEnd(comment) < arrowTokenIndex;
}
function handleArrowExpressionComments({
  comment,
  enclosingNode,
  followingNode,
  precedingNode,
  options,
}) {
  if (
    enclosingNode?.type !== "ArrowFunctionExpression" ||
    !followingNode ||
    !precedingNode
  ) {
    return false;
  }
  const isBeforeArrow = isCommentBeforeArrowFunctionExpressionArrow(
    comment,
    enclosingNode,
    options,
  );
  if (!isBeforeArrow) {
    addBlockOrNotComment(followingNode, comment);
    return true;
  }
  return false;
}
function getEnclosingAssignmentChainExpressionStatement(node, ancestors) {
  let child = node;
  for (const ancestor of ancestors) {
    if (
      (ancestor.type === "AssignmentExpression" && ancestor.right === child) ||
      (ancestor.type === "ArrowFunctionExpression" && ancestor.body === child)
    ) {
      child = ancestor;
      continue;
    }
    return ancestor.type === "ExpressionStatement" &&
      ancestor.expression === child
      ? ancestor
      : undefined;
  }
}
function handleSequenceExpressionLeadingComment({
  comment,
  enclosingNode,
  precedingNode,
  followingNode,
}) {
  if (
    !precedingNode &&
    enclosingNode?.type === "SequenceExpression" &&
    followingNode === enclosingNode.expressions[0]
  ) {
    addLeadingComment(enclosingNode, comment);
    return true;
  }
  return false;
}
function handleParenthesizedExpressionTrailingComment({
  comment,
  enclosingNode,
  precedingNode,
  followingNode,
  ancestors,
}) {
  if (!followingNode && enclosingNode && precedingNode) {
    if (
      enclosingNode.type === "ExpressionStatement" &&
      enclosingNode.expression === precedingNode
    ) {
      addTrailingComment(enclosingNode, comment);
      return true;
    }
    const isAssignment = precedingNode.type === "AssignmentExpression";
    if (
      (isAssignment &&
        enclosingNode.type === "AssignmentExpression" &&
        enclosingNode.right === precedingNode) ||
      (precedingNode.type === "ArrowFunctionExpression" &&
        enclosingNode.type === "ArrowFunctionExpression" &&
        enclosingNode.body === precedingNode)
    ) {
      const expressionStatement =
        getEnclosingAssignmentChainExpressionStatement(
          enclosingNode,
          ancestors.slice(1),
        );
      if (expressionStatement) {
        addTrailingComment(expressionStatement, comment);
        return true;
      }
    }
    const isSequence = precedingNode.type === "SequenceExpression";
    if (
      (isSequence || isAssignment) &&
      ((enclosingNode.type === "ArrowFunctionExpression" &&
        enclosingNode.body === precedingNode) ||
        (enclosingNode.type === "VariableDeclarator" &&
          enclosingNode.init === precedingNode) ||
        (enclosingNode.type === "ReturnStatement" &&
          enclosingNode.argument === precedingNode) ||
        (enclosingNode.type === "AssignmentExpression" &&
          enclosingNode.right === precedingNode))
    ) {
      addTrailingComment(
        isSequence ? precedingNode.expressions.at(-1) : precedingNode.right,
        comment,
      );
      return true;
    }
  }
  return false;
}
function handleUnionTypeLeadingComments(context) {
  const { followingNode, comment } = context;
  if (shouldAttachToUnionTypeFirstElement(followingNode, context)) {
    addLeadingComment(followingNode.types[0], comment);
    return true;
  }
  return false;
}
const isRealFunctionLikeNode = createTypeCheckFunction([
  "ArrowFunctionExpression",
  "FunctionExpression",
  "FunctionDeclaration",
  "ObjectMethod",
  "ClassMethod",
  "TSDeclareFunction",
  "TSCallSignatureDeclaration",
  "TSConstructSignatureDeclaration",
  "TSMethodSignature",
  "TSConstructorType",
  "TSFunctionType",
  "TSDeclareMethod",
  "HookDeclaration",
]);
const handleComments = {
  endOfLine: handleEndOfLineComment,
  ownLine: handleOwnLineComment,
  remaining: handleRemainingComment,
};
export default handleComments;
