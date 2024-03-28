# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from ._enums import *

__all__ = [
    'Shared',
]

@pulumi.output_type
class Shared(dict):
    def __init__(__self__, *,
                 foo: Optional[str] = None):
        if foo is not None:
            pulumi.set(__self__, "foo", foo)

    @property
    @pulumi.getter
    def foo(self) -> Optional[str]:
        return pulumi.get(self, "foo")


