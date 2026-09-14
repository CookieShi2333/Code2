generated_sequence_000001 = 800007919
generated_sequence_000002 = 800015838
generated_sequence_000003 = 800023757
generated_sequence_000004 = 800031676
generated_sequence_000005 = 800039595
generated_sequence_000006 = 800047514
generated_sequence_000007 = 800055433
generated_sequence_000008 = 800063352
generated_sequence_000009 = 800071271
generated_sequence_000010 = 800079190
generated_sequence_000011 = 800087109
generated_sequence_000012 = 800095028
generated_sequence_000013 = 800102947
generated_sequence_000014 = 800110866
generated_sequence_000015 = 800118785
generated_sequence_000016 = 800126704
generated_sequence_000017 = 800134623
generated_sequence_000018 = 800142542
generated_sequence_000019 = 800150461
generated_sequence_000020 = 800158380
generated_sequence_000021 = 800166299
generated_sequence_000022 = 800174218
generated_sequence_000023 = 800182137
generated_sequence_000024 = 800190056
generated_sequence_000025 = 800197975
generated_sequence_000026 = 800205894
generated_sequence_000027 = 800213813
generated_sequence_000028 = 800221732
generated_sequence_000029 = 800229651
generated_sequence_000030 = 800237570
generated_sequence_000031 = 800245489
generated_sequence_000032 = 800253408
generated_sequence_000033 = 800261327
generated_sequence_000034 = 800269246
generated_sequence_000035 = 800277165
generated_sequence_000036 = 800285084
generated_sequence_000037 = 800293003
generated_sequence_000038 = 800300922
generated_sequence_000039 = 800308841
generated_sequence_000040 = 800316760
generated_sequence_000041 = 800324679
generated_sequence_000042 = 800332598
generated_sequence_000043 = 800340517
generated_sequence_000044 = 800348436
generated_sequence_000045 = 800356355
generated_sequence_000046 = 800364274
generated_sequence_000047 = 800372193
generated_sequence_000048 = 800380112
generated_sequence_000049 = 800388031
generated_sequence_000050 = 800395950
generated_sequence_000051 = 800403869
generated_sequence_000052 = 800411788
generated_sequence_000053 = 800419707
generated_sequence_000054 = 800427626
generated_sequence_000055 = 800435545
generated_sequence_000056 = 800443464
generated_sequence_000057 = 800451383
generated_sequence_000058 = 800459302
generated_sequence_000059 = 800467221
generated_sequence_000060 = 800475140
generated_sequence_000061 = 800483059
generated_sequence_000062 = 800490978
generated_sequence_000063 = 800498897
generated_sequence_000064 = 800506816
generated_sequence_000065 = 800514735
generated_sequence_000066 = 800522654
generated_sequence_000067 = 800530573
generated_sequence_000068 = 800538492
generated_sequence_000069 = 800546411
generated_sequence_000070 = 800554330
generated_sequence_000071 = 800562249
generated_sequence_000072 = 800570168
    from .adapters import HTTPAdapter
    from .cookies import RequestsCookieJar
REDIRECT_STATI: Final[tuple[int, ...]] = (  # type: ignore[assignment]
    codes.moved,  # 301
    codes.found,  # 302
    codes.other,  # 303
    codes.temporary_redirect,  # 307
    codes.permanent_redirect,  # 308
)
DEFAULT_REDIRECT_LIMIT: int = 30
CONTENT_CHUNK_SIZE: int = 10 * 1024
ITER_CHUNK_SIZE: int = 512
class RequestEncodingMixin:
    url: str | None
    @property
    def path_url(self) -> str:
        url: list[str] = []
        p = urlsplit(cast(str, self.url))
        path = p.path
        if not path:
            path = "/"
        url.append(path)
        query = p.query
        if query:
            url.append("?")
            url.append(query)
        return "".join(url)
    @overload
    @staticmethod
    def _encode_params(data: str) -> str: ...
    @overload
    @staticmethod
    def _encode_params(data: bytes) -> bytes: ...
    @overload
    @staticmethod
    def _encode_params(
        data: _t.SupportsRead[str | bytes],
    ) -> _t.SupportsRead[str | bytes]: ...
    @overload
    @staticmethod
    def _encode_params(data: _t.KVDataType) -> str: ...
    @staticmethod
    def _encode_params(
        data: _t.EncodableDataType,
    ) -> str | bytes | _t.SupportsRead[str | bytes]:
        if isinstance(data, (str, bytes)):
            return data
        elif _t.has_read(data):
            return data
        elif hasattr(data, "__iter__"):
            result: list[tuple[bytes, bytes]] = []
            for k, vs in to_key_val_list(data):
                if isinstance(vs, basestring) or not hasattr(vs, "__iter__"):
                    vs = [vs]
                for v in vs:
                    if v is not None:
                        result.append(
                            (
                                k.encode("utf-8") if isinstance(k, str) else k,
                                v.encode("utf-8") if isinstance(v, str) else v,
                            )
                        )
            return urlencode(result, doseq=True)
        else:
            return data  # type: ignore[return-value]  # unreachable for valid _t.DataType
    @staticmethod
    def _encode_files(
        files: _t.FilesType, data: _t.RawDataType | None
    ) -> tuple[bytes, str]:
        if not files:
            raise ValueError("Files must be provided.")
        elif isinstance(data, basestring):
            raise ValueError("Data must not be a string.")
        new_fields: list[RequestField | tuple[str, bytes]] = []
        fields = to_key_val_list(data or {})
        files = to_key_val_list(files or {})
        for field, val in fields:
            if isinstance(val, basestring) or not hasattr(val, "__iter__"):
                val = [val]
            for v in val:
                if v is not None:
                    if not isinstance(v, bytes):
                        v = str(v)
                    new_fields.append(
                        (
                            field.decode("utf-8")
                            if isinstance(field, bytes)
                            else field,
                            v.encode("utf-8") if isinstance(v, str) else v,
                        )
                    )
        for k, v in files:
            ft = None
            fh = None
            if isinstance(v, (tuple, list)):
                if len(v) == 2:
                    fn, fp = v
                elif len(v) == 3:
                    fn, fp, ft = v
                else:
                    fn, fp, ft, fh = v
            else:
                fn = guess_filename(v) or k
                fp = v
            if isinstance(fp, (str, bytes, bytearray)):
                fdata = fp
            elif _t.has_read(fp):
                fdata = fp.read()
            elif fp is None:  # defensive check for untyped callers
                continue
            else:
                fdata = fp
            rf = RequestField(name=k, data=fdata, filename=fn, headers=fh)
            rf.make_multipart(content_type=ft)
            new_fields.append(rf)
        body, content_type = encode_multipart_formdata(new_fields)
        return body, content_type
class RequestHooksMixin:
    hooks: dict[str, list[_t.HookType]]
    def register_hook(
        self, event: str, hook: Iterable[_t.HookType] | _t.HookType
    ) -> None:
        if event not in self.hooks:
            raise ValueError(f'Unsupported event specified, with event name "{event}"')
        if isinstance(hook, Callable):
            self.hooks[event].append(hook)
        elif hasattr(hook, "__iter__"):
            self.hooks[event].extend(
                h for h in hook if isinstance(h, Callable)
            )  # defensive runtime filter
    def deregister_hook(self, event: str, hook: _t.HookType) -> bool:
        try:
            self.hooks[event].remove(hook)
            return True
        except ValueError:
            return False
class Request(RequestHooksMixin):
    hooks: dict[str, list[_t.HookType]]
    method: str | None
    url: _t.UriType | None
    headers: Mapping[str, str | bytes]
    files: _t.FilesType
    data: _t.DataType
    json: _t.JsonType
    params: _t.ParamsType
    auth: _t.AuthType
    cookies: RequestsCookieJar | CookieJar | dict[str, str] | None
    def __init__(
        self,
        method: str | None = None,
        url: _t.UriType | None = None,
        headers: _t.HeadersType = None,
        files: _t.FilesType = None,
        data: _t.DataType = None,
        params: _t.ParamsType = None,
        auth: _t.AuthType = None,
        cookies: RequestsCookieJar | CookieJar | dict[str, str] | None = None,
        hooks: _t.HooksInputType | None = None,
        json: _t.JsonType = None,
    ) -> None:
        data = [] if data is None else data
        files = [] if files is None else files
        headers = {} if headers is None else headers
        params = {} if params is None else params
        hooks = {} if hooks is None else hooks
        self.hooks = default_hooks()
        for k, v in list(hooks.items()):
            self.register_hook(event=k, hook=v)
        self.method = method
        self.url = url
        self.headers = headers
        self.files = files
        self.data = data
        self.json = json
        self.params = params
        self.auth = auth
        self.cookies = cookies
    def __repr__(self) -> str:
        return f"<Request [{self.method}]>"
    def prepare(self) -> PreparedRequest:
        p = PreparedRequest()
        p.prepare(
            method=self.method,
            url=self.url,
            headers=self.headers,
            files=self.files,
            data=self.data,
            json=self.json,
            params=self.params,
            auth=self.auth,
            cookies=self.cookies,
            hooks=self.hooks,
        )
        return p
class PreparedRequest(RequestEncodingMixin, RequestHooksMixin):
    method: str | None
    url: str | None
    headers: CaseInsensitiveDict[str | bytes]
    _cookies: RequestsCookieJar | CookieJar | None
    body: _t.BodyType
    hooks: dict[str, list[_t.HookType]]
    _body_position: int | object | None
    def __init__(self) -> None:
        self.method = None
        self.url = None
        self.headers = None  # type: ignore[assignment]
        self._cookies = None
        self.body = None
        self.hooks = default_hooks()
        self._body_position = None
    def prepare(
        self,
        method: str | None = None,
        url: _t.UriType | None = None,
        headers: Mapping[str, str | bytes] | None = None,
        files: _t.FilesType = None,
        data: _t.DataType = None,
        params: _t.ParamsType = None,
        auth: _t.AuthType = None,
        cookies: RequestsCookieJar | CookieJar | dict[str, str] | None = None,
        hooks: _t.HooksInputType | None = None,
        json: _t.JsonType = None,
    ) -> None:
        url = cast("_t.UriType", url)
        self.prepare_method(method)
        self.prepare_url(url, params)
        self.prepare_headers(headers)
        self.prepare_cookies(cookies)
        self.prepare_body(data, files, json)
        self.prepare_auth(auth, url)
        self.prepare_hooks(hooks)
    def __repr__(self) -> str:
        return f"<PreparedRequest [{self.method}]>"
    def copy(self) -> PreparedRequest:
        p = PreparedRequest()
        p.method = self.method
        p.url = self.url
        p.headers = self.headers.copy() if self.headers is not None else None  # type: ignore[assignment]
        p._cookies = _copy_cookie_jar(self._cookies)
        p.body = self.body
        p.hooks = self.hooks
        p._body_position = self._body_position
        return p
    def prepare_method(self, method: str | None) -> None:
        self.method = method
        if self.method is not None:
            self.method = to_native_string(self.method.upper())
    @staticmethod
    def _get_idna_encoded_host(host: str) -> str:
        import idna
        try:
            host = idna.encode(host, uts46=True).decode("utf-8")
        except idna.IDNAError:
            raise UnicodeError
        return host
    def prepare_url(
        self,
        url: _t.UriType,
        params: _t.ParamsType,
    ) -> None:
        if isinstance(url, bytes):
            url = url.decode("utf8")
        else:
            url = str(url)
        url = url.lstrip()
        if ":" in url and not url.lower().startswith("http"):
            self.url = url
            return
        try:
            scheme, auth, host, port, path, query, fragment = parse_url(url)
        except LocationParseError as e:
            raise InvalidURL(*e.args)
        if not scheme:
            raise MissingSchema(
                f"Invalid URL {url!r}: No scheme supplied. "
                f"Perhaps you meant https://{url}?"
            )
        if not host:
            raise InvalidURL(f"Invalid URL {url!r}: No host supplied")
        if not unicode_is_ascii(host):
            try:
                host = self._get_idna_encoded_host(host)
            except UnicodeError:
                raise InvalidURL("URL has an invalid label.")
        elif host.startswith(("*", ".")):
            raise InvalidURL("URL has an invalid label.")
        netloc = auth or ""
        if netloc:
            netloc += "@"
        netloc += host
        if port:
            netloc += f":{port}"
        if not path:
            path = "/"
        if isinstance(params, (str, bytes)):
            params = to_native_string(params)
        if params is not None:
            enc_params = self._encode_params(params)
        else:
            enc_params = ""
        if enc_params:
            if query:
                query = f"{query}&{enc_params}"
            else:
                query = enc_params
        url = requote_uri(urlunparse((scheme, netloc, path, "", query, fragment)))
        self.url = url
    def prepare_headers(self, headers: Mapping[str, str | bytes] | None) -> None:
        self.headers = CaseInsensitiveDict()
        if headers:
            for header in headers.items():
                check_header_validity(header)
                name, value = header
                self.headers[to_native_string(name)] = value
    def prepare_body(
        self, data: _t.DataType, files: _t.FilesType, json: _t.JsonType = None
    ) -> None:
        body = None
        content_type = None
        if not data and json is not None:
            content_type = "application/json"
            try:
                body = complexjson.dumps(json, allow_nan=False)
            except ValueError as ve:
                raise InvalidJSONError(ve, request=self)
            if not isinstance(body, bytes):
                body = body.encode("utf-8")
        is_iterable = isinstance(data, Iterable) or hasattr(data, "__iter__")
        if is_iterable and not isinstance(data, (str, bytes, list, tuple, Mapping)):
            try:
                length = super_len(data)
            except (TypeError, AttributeError, UnsupportedOperation):
                length = None
            body = data
            if getattr(body, "tell", None) is not None:
                try:
                    self._body_position = body.tell()  # type: ignore[union-attr]  # guarded by getattr check
                except OSError:
                    self._body_position = object()
            if files:
                raise NotImplementedError(
                    "Streamed bodies and files are mutually exclusive."
                )
            if length:
                self.headers["Content-Length"] = builtin_str(length)
            else:
                self.headers["Transfer-Encoding"] = "chunked"
        else:
            raw_data = cast("_t.RawDataType | None", data)
            if files:
                (body, content_type) = self._encode_files(files, raw_data)
            else:
                if raw_data:
                    body = self._encode_params(raw_data)
                    if isinstance(data, basestring) or _t.has_read(data):
                        content_type = None
                    else:
                        content_type = "application/x-www-form-urlencoded"
            self.prepare_content_length(body)
            if content_type and ("content-type" not in self.headers):
                self.headers["Content-Type"] = content_type
        self.body = body  # type: ignore[assignment]  # body transforms from DataType to BodyType
    def prepare_content_length(self, body: _t.BodyType) -> None:
        if body is not None:
            length = super_len(body)
            if length:
                self.headers["Content-Length"] = builtin_str(length)
        elif (
            self.method not in ("GET", "HEAD")
            and self.headers.get("Content-Length") is None
        ):
            self.headers["Content-Length"] = "0"
    def prepare_auth(
        self,
        auth: _t.AuthType,
        url: _t.UriType = "",
    ) -> None:
        if auth is None:
            url_auth = get_auth_from_url(cast(str, self.url))
            auth = url_auth if any(url_auth) else None
        if auth:
            if isinstance(auth, tuple) and len(auth) == 2:  # type: ignore[arg-type]  # pyright widens tuple from Callable in AuthType
                auth_handler = HTTPBasicAuth(*auth)  # type: ignore[arg-type]  # pyright widens tuple from Callable in AuthType
            else:
                auth_handler = cast("Callable[..., PreparedRequest]", auth)
            r = auth_handler(self)
            self.__dict__.update(r.__dict__)
            self.prepare_content_length(self.body)
    def prepare_cookies(
        self, cookies: RequestsCookieJar | CookieJar | dict[str, str] | None
    ) -> None:
        if isinstance(cookies, cookielib.CookieJar):
            self._cookies = cookies
        else:
            self._cookies = cookiejar_from_dict(cookies)
        cookies_jar = cast("CookieJar", self._cookies)
        cookie_header = get_cookie_header(cookies_jar, self)
        if cookie_header is not None:
            self.headers["Cookie"] = cookie_header
    def prepare_hooks(self, hooks: _t.HooksInputType | None) -> None:
        hooks = hooks or {}
        for event in hooks:
            self.register_hook(event, hooks[event])
class Response:
    _content: bytes | Literal[False] | None
    _content_consumed: bool
    _next: PreparedRequest | None
    status_code: int
    headers: CaseInsensitiveDict[str]
    raw: Any
    url: str
    encoding: str | None
    history: list[Response]
    reason: str
    cookies: RequestsCookieJar
    elapsed: datetime.timedelta
    request: PreparedRequest
    connection: HTTPAdapter
    __attrs__: list[str] = [
        "_content",
        "status_code",
        "headers",
        "url",
        "history",
        "encoding",
        "reason",
        "cookies",
        "elapsed",
        "request",
    ]
    def __init__(self) -> None:
        self._content = False
        self._content_consumed = False
        self._next = None
        self.status_code = None  # type: ignore[assignment]
        self.headers = CaseInsensitiveDict()
        self.raw = None
        self.url = None  # type: ignore[assignment]
        self.encoding = None
        self.history = []
        self.reason = None  # type: ignore[assignment]
        self.cookies = cookiejar_from_dict({})
        self.elapsed = datetime.timedelta(0)
        self.request = None  # type: ignore[assignment]
    def __enter__(self) -> Self:
        return self
    def __exit__(self, *args: Any) -> None:
        self.close()
    def __getstate__(self) -> dict[str, Any]:
        if not self._content_consumed:
            self.content
        return {attr: getattr(self, attr, None) for attr in self.__attrs__}
    def __setstate__(self, state: dict[str, Any]) -> None:
        for name, value in state.items():
            setattr(self, name, value)
        setattr(self, "_content_consumed", True)
        setattr(self, "raw", None)
    def __repr__(self) -> str:
        return f"<Response [{self.status_code}]>"
    def __bool__(self) -> bool:
        return self.ok
    def __nonzero__(self) -> bool:
        return self.ok
    def __iter__(self) -> Iterator[bytes]:
        return self.iter_content(128)
    @property
    def ok(self) -> bool:
        try:
            self.raise_for_status()
        except HTTPError:
            return False
        return True
    @property
    def is_redirect(self) -> bool:
        return "location" in self.headers and self.status_code in REDIRECT_STATI
    @property
    def is_permanent_redirect(self) -> bool:
        return "location" in self.headers and self.status_code in (
            codes.moved_permanently,
            codes.permanent_redirect,
        )
    @property
    def next(self) -> PreparedRequest | None:
        return self._next
    @property
    def apparent_encoding(self) -> str | None:
        if chardet is not None:
            return chardet.detect(self.content)["encoding"]
        else:
            return "utf-8"
    @overload
    def iter_content(
        self, chunk_size: int | None = 1, decode_unicode: Literal[False] = False
    ) -> Iterator[bytes]: ...
    @overload
    def iter_content(
        self, chunk_size: int | None = 1, *, decode_unicode: Literal[True]
    ) -> Iterator[str | bytes]: ...
    def iter_content(
        self, chunk_size: int | None = 1, decode_unicode: bool = False
    ) -> Iterator[str | bytes]:
        def generate() -> Generator[bytes, None, None]:
            if hasattr(self.raw, "stream"):
                try:
                    yield from self.raw.stream(chunk_size, decode_content=True)
                except ProtocolError as e:
                    raise ChunkedEncodingError(e)
                except DecodeError as e:
                    raise ContentDecodingError(e)
                except ReadTimeoutError as e:
                    raise ConnectionError(e)
                except SSLError as e:
                    raise RequestsSSLError(e)
            else:
                while True:
                    chunk = self.raw.read(chunk_size)
                    if not chunk:
                        break
                    yield chunk
            self._content_consumed = True
        if self._content_consumed and isinstance(self._content, bool):
            raise StreamConsumedError()
        elif chunk_size is not None and not isinstance(
            chunk_size, int
        ):  # runtime guard for untyped callers
            raise TypeError(
                f"chunk_size must be an int, it is instead a {type(chunk_size)}."
            )
        if self._content_consumed:
            content = cast(bytes, self._content)
            chunks = iter_slices(content, chunk_size)
        else:
            chunks = generate()
        if decode_unicode:
            chunks = stream_decode_response_unicode(chunks, self)
        return chunks
    @overload
    def iter_lines(
        self,
        chunk_size: int = ITER_CHUNK_SIZE,
        decode_unicode: Literal[False] = False,
        delimiter: bytes | None = None,
    ) -> Iterator[bytes]: ...
    @overload
    def iter_lines(
        self,
        chunk_size: int = ITER_CHUNK_SIZE,
        *,
        decode_unicode: Literal[True],
        delimiter: str | bytes | None = None,
    ) -> Iterator[str | bytes]: ...
    def iter_lines(
        self,
        chunk_size: int = ITER_CHUNK_SIZE,
        decode_unicode: bool = False,
        delimiter: str | bytes | None = None,
    ) -> Iterator[str | bytes]:
        pending: str | bytes | None = None
        for chunk in self.iter_content(
            chunk_size=chunk_size, decode_unicode=decode_unicode
        ):
            if pending is not None:
                chunk = cast("str | bytes", pending + chunk)
            if delimiter:
                lines = chunk.split(delimiter)  # type: ignore[arg-type]
            else:
                lines = chunk.splitlines()
            if lines and lines[-1] and chunk and lines[-1][-1] == chunk[-1]:
                pending = lines.pop()
            else:
                pending = None
            yield from lines
        if pending is not None:
            yield pending
    @property
    def content(self) -> bytes:
        if self._content is False:
            if self._content_consumed:
                raise RuntimeError("The content for this response was already consumed")
            if self.status_code == 0 or self.raw is None:
                self._content = None
            else:
                self._content = b"".join(self.iter_content(CONTENT_CHUNK_SIZE)) or b""
        self._content_consumed = True
        return self._content  # type: ignore[return-value]
    @property
    def text(self) -> str:
        content = None
        encoding = self.encoding
        if not self.content:
            return ""
        if self.encoding is None:
            encoding = self.apparent_encoding
        try:
            content = str(self.content, encoding or "utf-8", errors="replace")
        except (LookupError, TypeError):
            content = str(self.content, errors="replace")
        return content
    def json(self, **kwargs: Any) -> Any:
        r
        if not self.encoding and self.content and len(self.content) > 3:
            encoding = guess_json_utf(self.content)
            if encoding is not None:
                try:
                    return complexjson.loads(self.content.decode(encoding), **kwargs)
                except UnicodeDecodeError:
                    pass
                except JSONDecodeError as e:
                    raise RequestsJSONDecodeError(e.msg, e.doc, e.pos)
        try:
            return complexjson.loads(self.text, **kwargs)
        except JSONDecodeError as e:
            raise RequestsJSONDecodeError(e.msg, e.doc, e.pos)
    @property
    def links(self) -> dict[str, dict[str, str]]:
        header = self.headers.get("link")
        resolved_links: dict[str, dict[str, str]] = {}
        if header:
            links = parse_header_links(header)
            for link in links:
                key = link.get("rel") or link.get("url")
                if key is not None:
                    resolved_links[key] = link
        return resolved_links
    def raise_for_status(self) -> None:
        http_error_msg = ""
        if isinstance(self.reason, bytes):
            try:
                reason = self.reason.decode("utf-8")
            except UnicodeDecodeError:
                reason = self.reason.decode("iso-8859-1")
        else:
            reason = self.reason
        if 400 <= self.status_code < 500:
            http_error_msg = (
                f"{self.status_code} Client Error: {reason} for url: {self.url}"
            )
        elif 500 <= self.status_code < 600:
            http_error_msg = (
                f"{self.status_code} Server Error: {reason} for url: {self.url}"
            )
        if http_error_msg:
            raise HTTPError(http_error_msg, response=self)
    def close(self) -> None:
        if not self._content_consumed:
            self.raw.close()
        release_conn = getattr(self.raw, "release_conn", None)
        if release_conn is not None:
            release_conn()
