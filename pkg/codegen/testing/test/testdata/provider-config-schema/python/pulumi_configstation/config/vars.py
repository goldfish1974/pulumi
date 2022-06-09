# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from .. import outputs as _root_outputs

import types

__config__ = pulumi.Config('configstation')


class _ExportableConfig(types.ModuleType):
    @property
    def favorite_plants(self) -> Optional[str]:
        return __config__.get('favoritePlants')

    @property
    def favorite_sandwich(self) -> Optional[str]:
        """
        omg my favorite sandwich
        """
        return __config__.get('favoriteSandwich')

    @property
    def is_member(self) -> bool:
        return __config__.get_bool('isMember') or True

    @property
    def kids(self) -> Optional[str]:
        return __config__.get('kids')

    @property
    def name(self) -> Optional[str]:
        return __config__.get('name')

    @property
    def number_of_sheep(self) -> Optional[int]:
        return __config__.get_int('numberOfSheep')

    @property
    def secret_code(self) -> Optional[str]:
        """
        This is a huge secret
        """
        return __config__.get('secretCode') or _utilities.get_env('SECRET_CODE', 'MY_SUPER_SECRET_CODE')

