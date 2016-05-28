import {OnInit, Component} from '@angular/core'
import {ROUTER_DIRECTIVES} from '@angular/router';
import {RoleService} from "./services/roles.service";
import {Role} from "./contracts/role";

@Component({
    selector: 'wolvs-app',
    templateUrl: 'app/templates/newgame.html',
    directives: [ROUTER_DIRECTIVES]
})

export class NewGameComponent implements OnInit{
    constructor (private roleService: RoleService) {}

    ngOnInit() { this.getRoles() }

    errorMessage: string;
    roles: Role[];
    selectedRoles: Role[];

    private getRoles() {
        this.roleService.getRoles()
            .subscribe(
            roles =>
                this.roles = roles,
            error =>
                this.errorMessage = <any>error);
    }
}