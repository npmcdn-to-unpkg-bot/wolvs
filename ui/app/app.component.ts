import {Component} from '@angular/core'
import {Routes, ROUTER_DIRECTIVES, Router} from '@angular/router';
import {NewGameComponent} from './NewGame.component'
import {JoinGameComponent} from './JoinGame.component'
import {WelcomeComponent} from './Welcome.component'

@Component({
    selector: 'wolvs-app',
    directives: [ROUTER_DIRECTIVES],
    template: `<router-outlet></router-outlet>`
})

@Routes([
    {path: '/', component: WelcomeComponent},
    {path: '/NewGame', component: NewGameComponent},
    {path: '/JoinGame', component: JoinGameComponent}
])

export class AppComponent {
    constructor(private _router:Router){

    }
}