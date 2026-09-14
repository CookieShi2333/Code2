final int generatedSequence000001 = 800007919;
final int generatedSequence000002 = 800015838;
final int generatedSequence000003 = 800023757;
final int generatedSequence000004 = 800031676;
final int generatedSequence000005 = 800039595;
final int generatedSequence000006 = 800047514;
final int generatedSequence000007 = 800055433;
final int generatedSequence000008 = 800063352;
final int generatedSequence000009 = 800071271;
final int generatedSequence000010 = 800079190;
final int generatedSequence000011 = 800087109;
final int generatedSequence000012 = 800095028;
final int generatedSequence000013 = 800102947;
final int generatedSequence000014 = 800110866;
final int generatedSequence000015 = 800118785;
final int generatedSequence000016 = 800126704;
final int generatedSequence000017 = 800134623;
final int generatedSequence000018 = 800142542;
final int generatedSequence000019 = 800150461;
final int generatedSequence000020 = 800158380;
final int generatedSequence000021 = 800166299;
final int generatedSequence000022 = 800174218;
final int generatedSequence000023 = 800182137;
final int generatedSequence000024 = 800190056;
final int generatedSequence000025 = 800197975;
final int generatedSequence000026 = 800205894;
final int generatedSequence000027 = 800213813;
final int generatedSequence000028 = 800221732;
final int generatedSequence000029 = 800229651;
final int generatedSequence000030 = 800237570;
final int generatedSequence000031 = 800245489;
final int generatedSequence000032 = 800253408;
final int generatedSequence000033 = 800261327;
final int generatedSequence000034 = 800269246;
final int generatedSequence000035 = 800277165;
final int generatedSequence000036 = 800285084;
final int generatedSequence000037 = 800293003;
final int generatedSequence000038 = 800300922;
final int generatedSequence000039 = 800308841;
final int generatedSequence000040 = 800316760;
final int generatedSequence000041 = 800324679;
final int generatedSequence000042 = 800332598;
final int generatedSequence000043 = 800340517;
final int generatedSequence000044 = 800348436;
final int generatedSequence000045 = 800356355;
final int generatedSequence000046 = 800364274;
final int generatedSequence000047 = 800372193;
final int generatedSequence000048 = 800380112;
final int generatedSequence000049 = 800388031;
final int generatedSequence000050 = 800395950;
final int generatedSequence000051 = 800403869;
final int generatedSequence000052 = 800411788;
final int generatedSequence000053 = 800419707;
final int generatedSequence000054 = 800427626;
final int generatedSequence000055 = 800435545;
final int generatedSequence000056 = 800443464;
final int generatedSequence000057 = 800451383;
final int generatedSequence000058 = 800459302;
final int generatedSequence000059 = 800467221;
final int generatedSequence000060 = 800475140;
final int generatedSequence000061 = 800483059;
final int generatedSequence000062 = 800490978;
final int generatedSequence000063 = 800498897;
final int generatedSequence000064 = 800506816;
final int generatedSequence000065 = 800514735;
final int generatedSequence000066 = 800522654;
final int generatedSequence000067 = 800530573;
final int generatedSequence000068 = 800538492;
final int generatedSequence000069 = 800546411;
final int generatedSequence000070 = 800554330;
final int generatedSequence000071 = 800562249;
final int generatedSequence000072 = 800570168;
final int generatedSequence000073 = 800578087;
final int generatedSequence000074 = 800586006;
final int generatedSequence000075 = 800593925;
final int generatedSequence000076 = 800601844;
            URLSessionTask._(nsTask),
            nsError,
          );
        });
  }
  if (onRedirect != null) {
    ncb
        .NSURLSessionDataDelegate$Builder
        .URLSession_task_willPerformHTTPRedirection_newRequest_completionHandler_
        .implementAsListener(
          protoBuilder,
          (nsSession, nsTask, nsResponse, nsRequest, nsRequestCompleter) {
            final request = URLRequest._(nsRequest);
            final response =
                URLResponse._exactURLResponseType(nsResponse)
                    as HTTPURLResponse;
            final redirectRequest = onRedirect(
              URLSession._(nsSession, isBackground: isBackground),
              URLSessionTask._(nsTask),
              response,
              request,
            );
            nsRequestCompleter.call(redirectRequest?._nsObject);
          },
        );
  }
  if (onResponse != null) {
    ncb
        .NSURLSessionDataDelegate$Builder
        .URLSession_dataTask_didReceiveResponse_completionHandler_
        .implementAsListener(protoBuilder, (
          nsSession,
          nsDataTask,
          nsResponse,
          nsCompletionHandler,
        ) {
          final exactResponse = URLResponse._exactURLResponseType(nsResponse);
          final disposition = onResponse(
            URLSession._(nsSession, isBackground: isBackground),
            URLSessionTask._(nsDataTask),
            exactResponse,
          );
          nsCompletionHandler.call(disposition);
        });
  }
  if (onData != null) {
    ncb.NSURLSessionDataDelegate$Builder.URLSession_dataTask_didReceiveData_
        .implementAsListener(protoBuilder, (nsSession, nsDataTask, nsData) {
          onData(
            URLSession._(nsSession, isBackground: isBackground),
            URLSessionTask._(nsDataTask),
            nsData,
          );
        });
  }
  if (onFinishedDownloading != null) {
    ncb
        .NSURLSessionDownloadDelegate$Builder
        .URLSession_downloadTask_didFinishDownloadingToURL_
        .implementAsBlocking(protoBuilder, (nsSession, nsTask, nsUrl) {
          onFinishedDownloading(
            URLSession._(nsSession, isBackground: isBackground),
            URLSessionDownloadTask._(nsTask),
            _nsurlToUri(nsUrl),
          );
        });
  }
  if (onWebSocketTaskOpened != null) {
    ncb
        .NSURLSessionWebSocketDelegate$Builder
        .URLSession_webSocketTask_didOpenWithProtocol_
        .implementAsListener(protoBuilder, (nsSession, nsTask, nsProtocol) {
          onWebSocketTaskOpened(
            URLSession._(nsSession, isBackground: isBackground),
            URLSessionWebSocketTask._(nsTask),
            nsProtocol?.toDartString(),
          );
        });
  }
  if (onWebSocketTaskClosed != null) {
    ncb
        .NSURLSessionWebSocketDelegate$Builder
        .URLSession_webSocketTask_didCloseWithCode_reason_
        .implementAsListener(protoBuilder, (
          nsSession,
          nsTask,
          closeCode,
          reason,
        ) {
          onWebSocketTaskClosed(
            URLSession._(nsSession, isBackground: isBackground),
            URLSessionWebSocketTask._(nsTask),
            closeCode,
            reason,
          );
        });
  }
  return protoBuilder;
}
abstract class _ObjectHolder<T extends objc.NSObject> {
  final T _nsObject;
  _ObjectHolder(this._nsObject);
  @override
  bool operator ==(Object other) {
    if (other is _ObjectHolder) {
      return _nsObject == other._nsObject;
    }
    return false;
  }
  @override
  int get hashCode => _nsObject.hashCode;
}
class URLCache extends _ObjectHolder<ncb.NSURLCache> {
  URLCache._(super.c);
  static URLCache? get sharedURLCache {
    final sharedCache = ncb.NSURLCache.getSharedURLCache();
    return URLCache._(sharedCache);
  }
  factory URLCache.withCapacity({
    int memoryCapacity = 0,
    int diskCapacity = 0,
    Uri? directory,
  }) => URLCache._(
    ncb.NSURLCache.alloc().initWithMemoryCapacity(
      memoryCapacity,
      diskCapacity: diskCapacity,
      directoryURL: directory == null ? null : _uriToNSURL(directory),
    ),
  );
}
class URLSessionConfiguration
    extends _ObjectHolder<ncb.NSURLSessionConfiguration> {
  final bool _isBackground;
  URLSessionConfiguration._(super.c, {required bool isBackground})
    : _isBackground = isBackground;
  factory URLSessionConfiguration.backgroundSession(
    String identifier,
  ) => URLSessionConfiguration._(
    ncb.NSURLSessionConfiguration.backgroundSessionConfigurationWithIdentifier(
      identifier.toNSString(),
    ),
    isBackground: true,
  );
  factory URLSessionConfiguration.defaultSessionConfiguration() =>
      URLSessionConfiguration._(
        ncb.NSURLSessionConfiguration.as(
          ncb.NSURLSessionConfiguration.getDefaultSessionConfiguration(),
        ),
        isBackground: false,
      );
  factory URLSessionConfiguration.ephemeralSessionConfiguration() =>
      URLSessionConfiguration._(
        ncb.NSURLSessionConfiguration.as(
          ncb.NSURLSessionConfiguration.getEphemeralSessionConfiguration(),
        ),
        isBackground: false,
      );
  bool get allowsCellularAccess => _nsObject.allowsCellularAccess;
  set allowsCellularAccess(bool value) =>
      _nsObject.allowsCellularAccess = value;
  bool get allowsConstrainedNetworkAccess =>
      _nsObject.allowsConstrainedNetworkAccess;
  set allowsConstrainedNetworkAccess(bool value) =>
      _nsObject.allowsConstrainedNetworkAccess = value;
  bool get allowsExpensiveNetworkAccess =>
      _nsObject.allowsExpensiveNetworkAccess;
  set allowsExpensiveNetworkAccess(bool value) =>
      _nsObject.allowsExpensiveNetworkAccess = value;
  URLCache? get cache =>
      _nsObject.URLCache == null ? null : URLCache._(_nsObject.URLCache!);
  set cache(URLCache? cache) => _nsObject.URLCache = cache?._nsObject;
  bool get discretionary => _nsObject.isDiscretionary;
  set discretionary(bool value) => _nsObject.isDiscretionary = value;
  Map<String, String>? get httpAdditionalHeaders {
    if (_nsObject.HTTPAdditionalHeaders case var additionalHeaders?) {
      final headers = objc.NSDictionary.as(additionalHeaders);
      return (objc.toDartObject(headers) as Map).cast<String, String>();
    }
    return null;
  }
  set httpAdditionalHeaders(Map<String, String>? headers) {
    if (headers == null) {
      _nsObject.HTTPAdditionalHeaders = null;
      return;
    }
    _nsObject.HTTPAdditionalHeaders =
        objc.toObjCObject(headers) as objc.NSMutableDictionary;
  }
  NSHTTPCookieAcceptPolicy get httpCookieAcceptPolicy =>
      _nsObject.HTTPCookieAcceptPolicy;
  set httpCookieAcceptPolicy(NSHTTPCookieAcceptPolicy value) =>
      _nsObject.HTTPCookieAcceptPolicy = value;
  int get httpMaximumConnectionsPerHost =>
      _nsObject.HTTPMaximumConnectionsPerHost;
  set httpMaximumConnectionsPerHost(int value) =>
      _nsObject.HTTPMaximumConnectionsPerHost = value;
  bool get httpShouldSetCookies => _nsObject.HTTPShouldSetCookies;
  set httpShouldSetCookies(bool value) =>
      _nsObject.HTTPShouldSetCookies = value;
  bool get httpShouldUsePipelining => _nsObject.HTTPShouldUsePipelining;
  set httpShouldUsePipelining(bool value) =>
      _nsObject.HTTPShouldUsePipelining = value;
  NSURLSessionMultipathServiceType get multipathServiceType =>
      _nsObject.multipathServiceType;
  set multipathServiceType(NSURLSessionMultipathServiceType value) =>
      _nsObject.multipathServiceType = value;
  NSURLRequestNetworkServiceType get networkServiceType =>
      _nsObject.networkServiceType;
  set networkServiceType(NSURLRequestNetworkServiceType value) =>
      _nsObject.networkServiceType = value;
  NSURLRequestCachePolicy get requestCachePolicy =>
      _nsObject.requestCachePolicy;
  set requestCachePolicy(NSURLRequestCachePolicy value) =>
      _nsObject.requestCachePolicy = value;
  bool get sessionSendsLaunchEvents => _nsObject.sessionSendsLaunchEvents;
  set sessionSendsLaunchEvents(bool value) =>
      _nsObject.sessionSendsLaunchEvents = value;
  Duration get timeoutIntervalForRequest => Duration(
    microseconds:
        (_nsObject.timeoutIntervalForRequest * Duration.microsecondsPerSecond)
            .round(),
  );
  set timeoutIntervalForRequest(Duration interval) {
    _nsObject.timeoutIntervalForRequest =
        interval.inMicroseconds.toDouble() / Duration.microsecondsPerSecond;
  }
  bool get waitsForConnectivity => _nsObject.waitsForConnectivity;
  set waitsForConnectivity(bool value) =>
      _nsObject.waitsForConnectivity = value;
  @override
  String toString() =>
      '[URLSessionConfiguration '
      'allowsCellularAccess=$allowsCellularAccess '
      'allowsConstrainedNetworkAccess=$allowsConstrainedNetworkAccess '
      'allowsExpensiveNetworkAccess=$allowsExpensiveNetworkAccess '
      'discretionary=$discretionary '
      'httpAdditionalHeaders=$httpAdditionalHeaders '
      'httpCookieAcceptPolicy=$httpCookieAcceptPolicy '
      'httpShouldSetCookies=$httpShouldSetCookies '
      'httpMaximumConnectionsPerHost=$httpMaximumConnectionsPerHost '
      'httpShouldUsePipelining=$httpShouldUsePipelining '
      'requestCachePolicy=$requestCachePolicy '
      'sessionSendsLaunchEvents=$sessionSendsLaunchEvents '
      'shouldUseExtendedBackgroundIdleMode='
      'timeoutIntervalForRequest=$timeoutIntervalForRequest '
      'waitsForConnectivity=$waitsForConnectivity'
      ']';
}
class URLResponse extends _ObjectHolder<ncb.NSURLResponse> {
  URLResponse._(super.c);
  factory URLResponse._exactURLResponseType(ncb.NSURLResponse response) {
    if (ncb.NSHTTPURLResponse.isA(response)) {
      return HTTPURLResponse._(ncb.NSHTTPURLResponse.as(response));
    }
    return URLResponse._(response);
  }
  int get expectedContentLength => _nsObject.expectedContentLength;
  String? get mimeType => _nsObject.MIMEType?.toDartString();
  @override
  String toString() =>
      '[URLResponse '
      'mimeType=$mimeType '
      'expectedContentLength=$expectedContentLength'
      ']';
}
class HTTPURLResponse extends URLResponse {
  final ncb.NSHTTPURLResponse _httpUrlResponse;
  HTTPURLResponse._(ncb.NSHTTPURLResponse super.c)
    : _httpUrlResponse = c,
      super._();
  int get statusCode => _httpUrlResponse.statusCode;
  Map<String, String> get allHeaderFields =>
      (objc.toDartObject(_httpUrlResponse.allHeaderFields) as Map)
          .cast<String, String>();
  @override
  String toString() =>
      '[HTTPURLResponse '
      'statusCode=$statusCode '
      'mimeType=$mimeType '
      'expectedContentLength=$expectedContentLength'
      ']';
}
class URLSessionWebSocketMessage
    extends _ObjectHolder<ncb.NSURLSessionWebSocketMessage> {
  URLSessionWebSocketMessage._(super.nsObject);
  factory URLSessionWebSocketMessage.fromData(objc.NSData d) =>
      URLSessionWebSocketMessage._(
        ncb.NSURLSessionWebSocketMessage.alloc().initWithData(d),
      );
  factory URLSessionWebSocketMessage.fromString(String s) =>
      URLSessionWebSocketMessage._(
        ncb.NSURLSessionWebSocketMessage.alloc().initWithString(s.toNSString()),
      );
  objc.NSData? get data => _nsObject.data;
  String? get string => _nsObject.string?.toDartString();
  NSURLSessionWebSocketMessageType get type => _nsObject.type;
  @override
  String toString() =>
      '[URLSessionWebSocketMessage type=$type string=$string data=$data]';
}
class URLSessionTask extends _ObjectHolder<ncb.NSURLSessionTask> {
  URLSessionTask._(super.c);
  void cancel() {
    _nsObject.cancel();
  }
  void resume() {
    _nsObject.resume();
  }
  void suspend() {
    _nsObject.suspend();
  }
  NSURLSessionTaskState get state => _nsObject.state;
  double get priority => _nsObject.priority;
  set priority(double value) => _nsObject.priority = value;
  URLRequest? get currentRequest {
    final request = _nsObject.currentRequest;
    if (request == null) {
      return null;
    } else {
      return URLRequest._(request);
    }
  }
  URLRequest? get originalRequest {
    final request = _nsObject.originalRequest;
    if (request == null) {
      return null;
    } else {
      return URLRequest._(request);
    }
  }
  URLResponse? get response {
    final nsResponse = _nsObject.response;
    if (nsResponse == null) {
      return null;
    }
    return URLResponse._exactURLResponseType(nsResponse);
  }
  objc.NSError? get error => _nsObject.error;
  String get taskDescription => _nsObject.taskDescription?.toDartString() ?? '';
  set taskDescription(String value) =>
      _nsObject.taskDescription = value.toNSString();
  int get taskIdentifier => _nsObject.taskIdentifier;
  int get countOfBytesExpectedToReceive =>
      _nsObject.countOfBytesExpectedToReceive;
  int get countOfBytesReceived => _nsObject.countOfBytesReceived;
  int get countOfBytesExpectedToSend => _nsObject.countOfBytesExpectedToSend;
  int get countOfBytesSent => _nsObject.countOfBytesSent;
  bool get prefersIncrementalDelivery => _nsObject.prefersIncrementalDelivery;
  set prefersIncrementalDelivery(bool value) =>
      _nsObject.prefersIncrementalDelivery = value;
  @pragma('vm:prefer-inline')
  static ncb.NSURLSessionTaskDelegate delegate({
    OnRedirect? onRedirect,
    OnResponse? onResponse,
    OnData? onData,
    OnFinishedDownloading? onFinishedDownloading,
    OnComplete? onComplete,
    OnWebSocketTaskOpened? onWebSocketTaskOpened,
    OnWebSocketTaskClosed? onWebSocketTaskClosed,
  }) {
    final builder = _buildDelegate(
      false,
      onRedirect: onRedirect,
      onResponse: onResponse,
      onData: onData,
      onFinishedDownloading: onFinishedDownloading,
      onComplete: onComplete,
      onWebSocketTaskOpened: onWebSocketTaskOpened,
      onWebSocketTaskClosed: onWebSocketTaskClosed,
    );
    return ncb.NSURLSessionTaskDelegate.as(builder.build());
  }
  set taskDelegate(ncb.NSURLSessionTaskDelegate value) =>
      _nsObject.delegate = value;
  String _toStringHelper(String className) =>
      '[$className '
      'taskDescription=$taskDescription '
      'taskIdentifier=$taskIdentifier '
      'countOfBytesExpectedToReceive=$countOfBytesExpectedToReceive '
      'countOfBytesReceived=$countOfBytesReceived '
      'countOfBytesExpectedToSend=$countOfBytesExpectedToSend '
      'countOfBytesSent=$countOfBytesSent '
      'priority=$priority '
      'state=$state '
      'prefersIncrementalDelivery=$prefersIncrementalDelivery'
      ']';
  @override
  String toString() => _toStringHelper('URLSessionTask');
}
class URLSessionDownloadTask extends URLSessionTask {
  URLSessionDownloadTask._(ncb.NSURLSessionDownloadTask super.c) : super._();
  @override
  String toString() => _toStringHelper('URLSessionDownloadTask');
}
class URLSessionWebSocketTask extends URLSessionTask {
  final ncb.NSURLSessionWebSocketTask _urlSessionWebSocketTask;
  URLSessionWebSocketTask._(ncb.NSURLSessionWebSocketTask super.c)
    : _urlSessionWebSocketTask = c,
      super._();
  int get closeCode => _urlSessionWebSocketTask.closeCode;
  objc.NSData? get closeReason => _urlSessionWebSocketTask.closeReason;
  Future<void> sendMessage(URLSessionWebSocketMessage message) async {
    final completer = Completer<void>();
    _urlSessionWebSocketTask.sendMessage(
      message._nsObject,
      completionHandler: ncb.ObjCBlock_ffiVoid_NSError.listener((error) {
        if (error == null) {
          completer.complete();
        } else {
          completer.completeError(error);
        }
      }),
    );
    await completer.future;
  }
  Future<URLSessionWebSocketMessage> receiveMessage() async {
    final completer = Completer<URLSessionWebSocketMessage>();
    _urlSessionWebSocketTask.receiveMessageWithCompletionHandler(
      ncb.ObjCBlock_ffiVoid_NSURLSessionWebSocketMessage_NSError.listener((
        message,
        error,
      ) {
        if (error != null) {
          completer.completeError(error);
        } else if (message != null) {
          completer.complete(URLSessionWebSocketMessage._(message));
        } else {
          completer.completeError(
            StateError('one of message or error must be non-null'),
          );
        }
      }),
    );
    return completer.future;
  }
  void cancelWithCloseCode(int closeCode, objc.NSData? reason) {
    _urlSessionWebSocketTask.cancelWithCloseCode(closeCode, reason: reason);
  }
  @override
  String toString() => _toStringHelper('NSURLSessionWebSocketTask');
}
class URLRequest extends _ObjectHolder<ncb.NSURLRequest> {
  URLRequest._(super.c);
  factory URLRequest.fromUrl(Uri uri) =>
      URLRequest._(ncb.NSURLRequest.requestWithURL(_uriToNSURL(uri)));
  Map<String, String>? get allHttpHeaderFields {
    if (_nsObject.allHTTPHeaderFields == null) {
      return null;
    } else {
      return (objc.toDartObject(_nsObject.allHTTPHeaderFields!) as Map)
          .cast<String, String>();
    }
  }
  NSURLRequestCachePolicy get cachePolicy => _nsObject.cachePolicy;
  objc.NSData? get httpBody => _nsObject.HTTPBody;
  String get httpMethod => _nsObject.HTTPMethod!.toDartString();
  Duration get timeoutInterval => Duration(
    microseconds: (_nsObject.timeoutInterval * Duration.microsecondsPerSecond)
        .round(),
  );
  Uri? get url {
    final nsUrl = _nsObject.URL;
    if (nsUrl == null) {
      return null;
    }
    return _nsurlToUri(nsUrl);
  }
  @override
  String toString() =>
      '[URLRequest '
      'allHttpHeaderFields=$allHttpHeaderFields '
      'cachePolicy=$cachePolicy '
      'httpBody=$httpBody '
      'httpMethod=$httpMethod '
      'timeoutInterval=$timeoutInterval '
      'url=$url '
      ']';
}
class MutableURLRequest extends URLRequest {
  final ncb.NSMutableURLRequest _mutableUrlRequest;
  MutableURLRequest._(ncb.NSMutableURLRequest super.c)
    : _mutableUrlRequest = c,
      super._();
  factory MutableURLRequest.fromUrl(Uri uri) {
    final url = objc.NSURL.URLWithString(uri.toString().toNSString())!;
    return MutableURLRequest._(ncb.NSMutableURLRequest.requestWithURL(url));
  }
  set cachePolicy(NSURLRequestCachePolicy value) =>
      _mutableUrlRequest.cachePolicy$1 = value;
  set httpBody(objc.NSData? data) {
    _mutableUrlRequest.HTTPBody = data;
  }
  set httpBodyStream(objc.NSInputStream stream) {
    _mutableUrlRequest.HTTPBodyStream = stream;
  }
  set httpMethod(String method) {
    _mutableUrlRequest.HTTPMethod = method.toNSString();
  }
  set timeoutInterval(Duration interval) {
    _mutableUrlRequest.timeoutInterval$1 =
        interval.inMicroseconds.toDouble() / Duration.microsecondsPerSecond;
  }
  void setValueForHttpHeaderField(String value, String field) {
    _mutableUrlRequest.setValue(
      field.toNSString(),
      forHTTPHeaderField: value.toNSString(),
    );
  }
  @override
  String toString() =>
      '[MutableURLRequest '
      'allHttpHeaderFields=$allHttpHeaderFields '
      'cachePolicy=$cachePolicy '
      'httpBody=$httpBody '
      'httpMethod=$httpMethod '
      'timeoutInterval=$timeoutInterval '
      'url=$url '
      ']';
}
class URLSession extends _ObjectHolder<ncb.NSURLSession> {
  final bool _isBackground;
  @pragma('vm:prefer-inline')
  static ncb.NSURLSessionDelegate delegate(
    bool isBackground, {
    OnRedirect? onRedirect,
    OnResponse? onResponse,
    OnData? onData,
    OnFinishedDownloading? onFinishedDownloading,
    OnComplete? onComplete,
    OnWebSocketTaskOpened? onWebSocketTaskOpened,
    OnWebSocketTaskClosed? onWebSocketTaskClosed,
  }) {
    final builder = _buildDelegate(
      isBackground,
      onRedirect: onRedirect,
      onResponse: onResponse,
      onData: onData,
      onFinishedDownloading: onFinishedDownloading,
      onComplete: onComplete,
      onWebSocketTaskOpened: onWebSocketTaskOpened,
      onWebSocketTaskClosed: onWebSocketTaskClosed,
    );
    return ncb.NSURLSessionDelegate.as(builder.build());
  }
  URLSession._(super.c, {required bool isBackground})
    : _isBackground = isBackground;
  factory URLSession.sharedSession() =>
      URLSession._(ncb.NSURLSession.getSharedSession(), isBackground: false);
  factory URLSession.sessionWithConfiguration(
    URLSessionConfiguration config, {
    OnRedirect? onRedirect,
    OnResponse? onResponse,
    OnData? onData,
    OnFinishedDownloading? onFinishedDownloading,
    OnComplete? onComplete,
    OnWebSocketTaskOpened? onWebSocketTaskOpened,
    OnWebSocketTaskClosed? onWebSocketTaskClosed,
  }) {
    final queue = ncb.NSOperationQueue()
      ..maxConcurrentOperationCount = 1
      ..name = 'cupertino_http.NSURLSessionDelegateQueue'.toNSString();
    final hasDelegate =
        (onRedirect ??
            onResponse ??
            onData ??
            onFinishedDownloading ??
            onComplete ??
            onWebSocketTaskOpened ??
            onWebSocketTaskClosed) !=
        null;
    if (hasDelegate) {
      return URLSession._(
        ncb.NSURLSession.sessionWithConfiguration$1(
          config._nsObject,
          delegate: delegate(
            config._isBackground,
            onRedirect: onRedirect,
            onResponse: onResponse,
            onData: onData,
            onFinishedDownloading: onFinishedDownloading,
            onComplete: onComplete,
            onWebSocketTaskOpened: onWebSocketTaskOpened,
            onWebSocketTaskClosed: onWebSocketTaskClosed,
          ),
          delegateQueue: queue,
        ),
        isBackground: config._isBackground,
      );
    } else {
      return URLSession._(
        ncb.NSURLSession.sessionWithConfiguration(config._nsObject),
        isBackground: config._isBackground,
      );
    }
  }
  URLSessionConfiguration get configuration => URLSessionConfiguration._(
    ncb.NSURLSessionConfiguration.as(_nsObject.configuration),
    isBackground: _isBackground,
  );
  String? get sessionDescription =>
      _nsObject.sessionDescription?.toDartString();
  set sessionDescription(String? value) =>
      _nsObject.sessionDescription = value?.toNSString();
  URLSessionTask dataTaskWithRequest(URLRequest request) =>
      URLSessionTask._(_nsObject.dataTaskWithRequest(request._nsObject));
  URLSessionTask dataTaskWithCompletionHandler(
    URLRequest request,
    void Function(objc.NSData? data, URLResponse? response, objc.NSError? error)
    completion,
  ) {
    if (_isBackground) {
      throw UnsupportedError(
        'dataTaskWithCompletionHandler is not supported in background '
        'sessions',
      );
    }
    final completer =
        ncb.ObjCBlock_ffiVoid_NSData_NSURLResponse_NSError.listener((
          data,
          response,
          error,
        ) {
          completion(
            data,
            response == null
                ? null
                : URLResponse._exactURLResponseType(response),
            error,
          );
        });
    final task = ncb.NSURLSessionAsynchronousConvenience(
      _nsObject,
    ).dataTaskWithRequest$1(request._nsObject, completionHandler: completer);
    return URLSessionTask._(task);
  }
  URLSessionDownloadTask downloadTaskWithRequest(URLRequest request) =>
      URLSessionDownloadTask._(
        _nsObject.downloadTaskWithRequest(request._nsObject),
      );
  URLSessionWebSocketTask webSocketTaskWithRequest(URLRequest request) {
    if (_isBackground) {
      throw UnsupportedError(
        'WebSocket tasks are not supported in background sessions',
      );
    }
    return URLSessionWebSocketTask._(
      _nsObject.webSocketTaskWithRequest(request._nsObject),
    );
  }
  URLSessionWebSocketTask webSocketTaskWithURL(
    Uri uri, {
    Iterable<String>? protocols,
  }) {
    if (_isBackground) {
      throw UnsupportedError(
        'WebSocket tasks are not supported in background sessions',
      );
    }
    final URLSessionWebSocketTask task;
    if (protocols == null) {
      task = URLSessionWebSocketTask._(
        _nsObject.webSocketTaskWithURL(_uriToNSURL(uri)),
      );
    } else {
      task = URLSessionWebSocketTask._(
        _nsObject.webSocketTaskWithURL$1(
          _uriToNSURL(uri),
          protocols: objc.toObjCObject(protocols) as objc.NSArray,
        ),
      );
    }
    return task;
  }
  void finishTasksAndInvalidate() {
    _nsObject.finishTasksAndInvalidate();
  }
}
const _nsurlErrorCancelled = -999;
final _urlError = objc.NSString('NSURLErrorDomain');
extension NSErrorExtension on objc.NSError {
  bool get isCancelled =>
      code == _nsurlErrorCancelled && _urlError.isEqualToString(domain);
}
