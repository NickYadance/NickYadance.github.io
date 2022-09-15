---
title: "Google Oauth"
description: ""
lead: ""
date: 2022-09-12T18:10:35+08:00
lastmod: 2022-09-12T18:10:35+08:00
draft: false
images: []
menu:
  docs:
    parent: ""
    identifier: "google_oauth-94bda30947ed48b2c444533b24c06a23"
weight: 999
toc: true
---
1. Implement `AirflowSecurityManager`, and put it under any PYTHONPATH in Airflow.  

```python
import logging
import os
from typing import Any, List, Union, Dict

from airflow.www.fab_security.sqla.models import Permission
from airflow.www.security import AirflowSecurityManager

log = logging.getLogger(__name__)
log.setLevel(os.getenv("AIRFLOW__LOGGING__FAB_LOGGING_LEVEL", "INFO"))
FAB_ADMIN_ROLE = "Admin"
FAB_OP_ROLE = "Op"
FAB_VIEWER_ROLE = "Viewer"
FAB_PUBLIC_ROLE = "Public"


class GoogleAuthorizer(AirflowSecurityManager):
    def get_role_permissions_from_db(self, role_id: int) -> List[Permission]:
        pass

    def get_oauth_user_info(
            self, provider: str, resp: Any
    ) -> Dict[str, Union[str, List[str]]]:
        authorized_hd = 'shopee.com'
        me = self.appbuilder.sm.oauth_remotes[provider].get("userinfo")
        data = me.json()
        log.info('google oauth user: %s', data)
        if data['hd'] != authorized_hd:
            return {"username": "guest", "role_keys": [FAB_PUBLIC_ROLE]}
        else:
            return {
                "username": data.get("name", ""),
                "role_keys": [FAB_OP_ROLE],
                "first_name": data.get("given_name", ""),
                "last_name": data.get("family_name", ""),
                "email": data.get("email", "")
            }
```

2. Update `$AIRFLOW_HOME/webserver_config.py`. Note that `FAB_SECURITY_MANAGER_CLASS` is the full package name of the SecurityManager in step1.
```python
import os

from flask_appbuilder.const import AUTH_OAUTH

basedir = os.path.abspath(os.path.dirname(__file__))

WTF_CSRF_ENABLED = True
AUTH_TYPE = AUTH_OAUTH
AUTH_ROLES_SYNC_AT_LOGIN = True  # Checks roles on every login
AUTH_USER_REGISTRATION = (
    True  # allow users who are not already in the FAB DB to register
)
AUTH_ROLES_MAPPING = {
    "Viewer": ["Viewer"],
    "Admin": ["Admin"],
    "Op": ["Op"],
}
FAB_SECURITY_MANAGER_CLASS = "security_manager.GoogleAuthorizer"
GOOGLE_KEY = ''
GOOGLE_SECRET_KEY = ''
OAUTH_PROVIDERS = [{
    'name': 'google',
    'token_key': 'access_token',
    'icon': 'fa-google',
    'remote_app': {
        'api_base_url': 'https://www.googleapis.com/oauth2/v2/',
        'client_kwargs': {
            'scope': 'email profile',
            "cookie_policy": "single_host_origin"
        },
        'access_token_url': 'https://accounts.google.com/o/oauth2/token',
        'authorize_url': 'https://accounts.google.com/o/oauth2/auth',
        'request_token_url': None,
        'client_id': GOOGLE_KEY,
        'client_secret': GOOGLE_SECRET_KEY,
    }
}]
```
