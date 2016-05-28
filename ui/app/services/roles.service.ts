import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';
import { Role } from '../contracts/role'
import 'rxjs/add/operator/map'
import 'rxjs/add/operator/catch'
import { Observable }     from 'rxjs/Observable';

@Injectable()
export class RoleService {
    constructor(private http: Http) {}

    private rolesUrl = 'http://localhost:8080/roles';

    getRoles (): Observable<Role[]> {
        return this.http.get(this.rolesUrl)
                        .map(this.extractData)
                        .catch(this.handleError);
    }

    private extractData(res: Response) {
        let body = res.json();
        return Array.from(body) || [];
    }

    private handleError (error: any) {
        // TODO: Choose a logging platform and log to it.
        let errMsg = (error.message) ? error.message :
            error.status ? `${error.status} - ${error.statusText}` : 'Server error';
        console.error(errMsg); // log to console for now
        return Observable.throw(errMsg);
    }

}