from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ExecuteActionsRequest(_message.Message):
    __slots__ = ("task_name", "path", "context")
    TASK_NAME_FIELD_NUMBER: _ClassVar[int]
    PATH_FIELD_NUMBER: _ClassVar[int]
    CONTEXT_FIELD_NUMBER: _ClassVar[int]
    task_name: str
    path: str
    context: Context
    def __init__(self, task_name: _Optional[str] = ..., path: _Optional[str] = ..., context: _Optional[_Union[Context, _Mapping]] = ...) -> None: ...

class ExecuteActionsResponse(_message.Message):
    __slots__ = ("error",)
    ERROR_FIELD_NUMBER: _ClassVar[int]
    error: str
    def __init__(self, error: _Optional[str] = ...) -> None: ...

class ExecuteFiltersRequest(_message.Message):
    __slots__ = ("task_name", "context")
    TASK_NAME_FIELD_NUMBER: _ClassVar[int]
    CONTEXT_FIELD_NUMBER: _ClassVar[int]
    task_name: str
    context: Context
    def __init__(self, task_name: _Optional[str] = ..., context: _Optional[_Union[Context, _Mapping]] = ...) -> None: ...

class ExecuteFiltersResponse(_message.Message):
    __slots__ = ("match", "error")
    MATCH_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    match: bool
    error: str
    def __init__(self, match: bool = ..., error: _Optional[str] = ...) -> None: ...

class ListTasksRequest(_message.Message):
    __slots__ = ("custom_config",)
    CUSTOM_CONFIG_FIELD_NUMBER: _ClassVar[int]
    custom_config: bytes
    def __init__(self, custom_config: _Optional[bytes] = ...) -> None: ...

class ListTasksResponse(_message.Message):
    __slots__ = ("tasks", "error")
    TASKS_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    tasks: _containers.RepeatedCompositeFieldContainer[Task]
    error: str
    def __init__(self, tasks: _Optional[_Iterable[_Union[Task, _Mapping]]] = ..., error: _Optional[str] = ...) -> None: ...

class Action(_message.Message):
    __slots__ = ("file",)
    FILE_FIELD_NUMBER: _ClassVar[int]
    file: ActionFile
    def __init__(self, file: _Optional[_Union[ActionFile, _Mapping]] = ...) -> None: ...

class ActionFile(_message.Message):
    __slots__ = ("content", "mode", "path", "state")
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    MODE_FIELD_NUMBER: _ClassVar[int]
    PATH_FIELD_NUMBER: _ClassVar[int]
    STATE_FIELD_NUMBER: _ClassVar[int]
    content: str
    mode: str
    path: str
    state: str
    def __init__(self, content: _Optional[str] = ..., mode: _Optional[str] = ..., path: _Optional[str] = ..., state: _Optional[str] = ...) -> None: ...

class Context(_message.Message):
    __slots__ = ("repository",)
    REPOSITORY_FIELD_NUMBER: _ClassVar[int]
    repository: Repository
    def __init__(self, repository: _Optional[_Union[Repository, _Mapping]] = ...) -> None: ...

class Repository(_message.Message):
    __slots__ = ("full_name", "clone_url_http", "clone_url_ssh", "web_url")
    FULL_NAME_FIELD_NUMBER: _ClassVar[int]
    CLONE_URL_HTTP_FIELD_NUMBER: _ClassVar[int]
    CLONE_URL_SSH_FIELD_NUMBER: _ClassVar[int]
    WEB_URL_FIELD_NUMBER: _ClassVar[int]
    full_name: str
    clone_url_http: str
    clone_url_ssh: str
    web_url: str
    def __init__(self, full_name: _Optional[str] = ..., clone_url_http: _Optional[str] = ..., clone_url_ssh: _Optional[str] = ..., web_url: _Optional[str] = ...) -> None: ...

class Filter(_message.Message):
    __slots__ = ("file", "line_in_file", "repository_name", "reverse")
    FILE_FIELD_NUMBER: _ClassVar[int]
    LINE_IN_FILE_FIELD_NUMBER: _ClassVar[int]
    REPOSITORY_NAME_FIELD_NUMBER: _ClassVar[int]
    REVERSE_FIELD_NUMBER: _ClassVar[int]
    file: FilterFile
    line_in_file: FilterLineInFile
    repository_name: FilterRepositoryName
    reverse: bool
    def __init__(self, file: _Optional[_Union[FilterFile, _Mapping]] = ..., line_in_file: _Optional[_Union[FilterLineInFile, _Mapping]] = ..., repository_name: _Optional[_Union[FilterRepositoryName, _Mapping]] = ..., reverse: bool = ...) -> None: ...

class FilterLineInFile(_message.Message):
    __slots__ = ("file", "search")
    FILE_FIELD_NUMBER: _ClassVar[int]
    SEARCH_FIELD_NUMBER: _ClassVar[int]
    file: str
    search: str
    def __init__(self, file: _Optional[str] = ..., search: _Optional[str] = ...) -> None: ...

class FilterFile(_message.Message):
    __slots__ = ("file",)
    FILE_FIELD_NUMBER: _ClassVar[int]
    file: str
    def __init__(self, file: _Optional[str] = ...) -> None: ...

class FilterRepositoryName(_message.Message):
    __slots__ = ("names",)
    NAMES_FIELD_NUMBER: _ClassVar[int]
    names: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, names: _Optional[_Iterable[str]] = ...) -> None: ...

class OnPrClosedRequest(_message.Message):
    __slots__ = ("task_name", "context")
    TASK_NAME_FIELD_NUMBER: _ClassVar[int]
    CONTEXT_FIELD_NUMBER: _ClassVar[int]
    task_name: str
    context: Context
    def __init__(self, task_name: _Optional[str] = ..., context: _Optional[_Union[Context, _Mapping]] = ...) -> None: ...

class OnPrClosedResponse(_message.Message):
    __slots__ = ("error",)
    ERROR_FIELD_NUMBER: _ClassVar[int]
    error: str
    def __init__(self, error: _Optional[str] = ...) -> None: ...

class OnPrCreatedRequest(_message.Message):
    __slots__ = ("task_name", "context")
    TASK_NAME_FIELD_NUMBER: _ClassVar[int]
    CONTEXT_FIELD_NUMBER: _ClassVar[int]
    task_name: str
    context: Context
    def __init__(self, task_name: _Optional[str] = ..., context: _Optional[_Union[Context, _Mapping]] = ...) -> None: ...

class OnPrCreatedResponse(_message.Message):
    __slots__ = ("error",)
    ERROR_FIELD_NUMBER: _ClassVar[int]
    error: str
    def __init__(self, error: _Optional[str] = ...) -> None: ...

class OnPrMergedRequest(_message.Message):
    __slots__ = ("task_name", "context")
    TASK_NAME_FIELD_NUMBER: _ClassVar[int]
    CONTEXT_FIELD_NUMBER: _ClassVar[int]
    task_name: str
    context: Context
    def __init__(self, task_name: _Optional[str] = ..., context: _Optional[_Union[Context, _Mapping]] = ...) -> None: ...

class OnPrMergedResponse(_message.Message):
    __slots__ = ("error",)
    ERROR_FIELD_NUMBER: _ClassVar[int]
    error: str
    def __init__(self, error: _Optional[str] = ...) -> None: ...

class Task(_message.Message):
    __slots__ = ("name", "auto_merge", "auto_merge_after_seconds", "branch_name", "change_limit", "commit_message", "create_only", "disabled", "keep_branch_after_merge", "labels", "merge_once", "pr_body", "pr_title", "filters", "actions")
    NAME_FIELD_NUMBER: _ClassVar[int]
    AUTO_MERGE_FIELD_NUMBER: _ClassVar[int]
    AUTO_MERGE_AFTER_SECONDS_FIELD_NUMBER: _ClassVar[int]
    BRANCH_NAME_FIELD_NUMBER: _ClassVar[int]
    CHANGE_LIMIT_FIELD_NUMBER: _ClassVar[int]
    COMMIT_MESSAGE_FIELD_NUMBER: _ClassVar[int]
    CREATE_ONLY_FIELD_NUMBER: _ClassVar[int]
    DISABLED_FIELD_NUMBER: _ClassVar[int]
    KEEP_BRANCH_AFTER_MERGE_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    MERGE_ONCE_FIELD_NUMBER: _ClassVar[int]
    PR_BODY_FIELD_NUMBER: _ClassVar[int]
    PR_TITLE_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    ACTIONS_FIELD_NUMBER: _ClassVar[int]
    name: str
    auto_merge: bool
    auto_merge_after_seconds: int
    branch_name: str
    change_limit: int
    commit_message: str
    create_only: bool
    disabled: bool
    keep_branch_after_merge: bool
    labels: _containers.RepeatedScalarFieldContainer[str]
    merge_once: bool
    pr_body: str
    pr_title: str
    filters: _containers.RepeatedCompositeFieldContainer[Filter]
    actions: _containers.RepeatedCompositeFieldContainer[Action]
    def __init__(self, name: _Optional[str] = ..., auto_merge: bool = ..., auto_merge_after_seconds: _Optional[int] = ..., branch_name: _Optional[str] = ..., change_limit: _Optional[int] = ..., commit_message: _Optional[str] = ..., create_only: bool = ..., disabled: bool = ..., keep_branch_after_merge: bool = ..., labels: _Optional[_Iterable[str]] = ..., merge_once: bool = ..., pr_body: _Optional[str] = ..., pr_title: _Optional[str] = ..., filters: _Optional[_Iterable[_Union[Filter, _Mapping]]] = ..., actions: _Optional[_Iterable[_Union[Action, _Mapping]]] = ...) -> None: ...
