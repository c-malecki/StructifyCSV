export namespace main {
	
	export class SelectFileRes {
	    headers: string[];
	    error: boolean;
	
	    static createFrom(source: any = {}) {
	        return new SelectFileRes(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.headers = source["headers"];
	        this.error = source["error"];
	    }
	}

}

