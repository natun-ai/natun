# -*- coding: utf-8 -*-
#  Copyright (c) 2022 RaptorML authors.
#
#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.

from typing import Dict, Union

from labsdk.raptor import SourceProductionConfig
from labsdk.raptor.types import SecretKeyRef


class RestConfig(SourceProductionConfig):
    @classmethod
    def kind(cls) -> str:
        return 'rest'

    def configurable_envs(self) -> Dict[str, str]:
        return {}

    def config(self) -> Dict[str, Union[str, SecretKeyRef]]:
        return {
            'url': '<http://localhost:5000/users/{key:user_id}/{keys}>',
            'method': 'GET',
            'body': '',
            'headers': ''
        }
