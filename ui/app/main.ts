import {bootstrap}    from '@angular/platform-browser-dynamic';
import {ROUTER_PROVIDERS} from '@angular/router';
import {HTTP_PROVIDERS} from '@angular/http';
import {AppComponent} from './app.component';
import {RoleService} from "./services/roles.service";

bootstrap(AppComponent, [ROUTER_PROVIDERS, HTTP_PROVIDERS, RoleService]);