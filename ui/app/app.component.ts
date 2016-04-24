import {Component} from 'angular2/core'
import {RouteConfig, ROUTER_DIRECTIVES} from 'angular2/router';
import {NewGameComponent} from './NewGame.component'
import {JoinGameComponent} from './JoinGame.component'
import {WelcomeComponent} from './Welcome.component'

@Component({
    selector: 'wolvs-app',
    directives: [ROUTER_DIRECTIVES],
    template: `<router-outlet></router-outlet>`
})

@RouteConfig([
    {path: '/', name:'Welcome', component: WelcomeComponent},
    {path: '/new', name: 'NewGame', component: NewGameComponent},
    {path: '/join', name: 'JoinGame', component: JoinGameComponent}
])

export class AppComponent { }