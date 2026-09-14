library;
import 'dart:async';
import 'package:objective_c/objective_c.dart' as objc;
import 'native_cupertino_bindings.dart' as ncb;
import 'native_cupertino_bindings.dart'
    show
        NSHTTPCookieAcceptPolicy,
        NSURLRequestCachePolicy,
        NSURLRequestNetworkServiceType,
        NSURLSessionMultipathServiceType,
        NSURLSessionResponseDisposition,
        NSURLSessionTaskState,
        NSURLSessionWebSocketMessageType;
export 'native_cupertino_bindings.dart'
    show
        NSHTTPCookieAcceptPolicy,
        NSURLRequestCachePolicy,
        NSURLRequestNetworkServiceType,
        NSURLSessionMultipathServiceType,
        NSURLSessionResponseDisposition,
        NSURLSessionTaskState,
        NSURLSessionWebSocketCloseCode,
        NSURLSessionWebSocketMessageType;
objc.NSURL _uriToNSURL(Uri uri) =>
    objc.NSURL.URLWithString(uri.toString().toNSString())!;
Uri _nsurlToUri(objc.NSURL url) =>
    Uri.parse(url.absoluteString!.toDartString());
typedef OnRedirect =
    URLRequest? Function(
      URLSession session,
      URLSessionTask task,
      HTTPURLResponse response,
      URLRequest newRequest,
    );
typedef OnResponse =
    NSURLSessionResponseDisposition Function(
      URLSession session,
      URLSessionTask task,
      URLResponse response,
    );
typedef OnData =
    void Function(URLSession session, URLSessionTask task, objc.NSData data);
typedef OnFinishedDownloading =
    void Function(URLSession session, URLSessionDownloadTask task, Uri uri);
typedef OnComplete =
    void Function(URLSession session, URLSessionTask task, objc.NSError? error);
typedef OnWebSocketTaskOpened =
    void Function(
      URLSession session,
      URLSessionWebSocketTask task,
      String? protocol,
    );
typedef OnWebSocketTaskClosed =
    void Function(
      URLSession session,
      URLSessionWebSocketTask task,
      int closeCode,
      objc.NSData? reason,
    );
@pragma('vm:prefer-inline')
objc.ObjCProtocolBuilder _buildDelegate(
  bool isBackground, {
  OnRedirect? onRedirect,
  OnResponse? onResponse,
  OnData? onData,
  OnFinishedDownloading? onFinishedDownloading,
  OnComplete? onComplete,
  OnWebSocketTaskOpened? onWebSocketTaskOpened,
  OnWebSocketTaskClosed? onWebSocketTaskClosed,
}) {
  final protoBuilder = objc.ObjCProtocolBuilder();
  if (onComplete != null) {
    ncb.NSURLSessionDataDelegate$Builder.URLSession_task_didCompleteWithError_
        .implementAsListener(protoBuilder, (nsSession, nsTask, nsError) {
          onComplete(
            URLSession._(nsSession, isBackground: isBackground),
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
