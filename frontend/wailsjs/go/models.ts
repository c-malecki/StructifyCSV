export namespace core {
	
	export class ImportSchemaRes {
	    schema?: entity.JsonSchema;
	    error: any;
	
	    static createFrom(source: any = {}) {
	        return new ImportSchemaRes(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.schema = this.convertValues(source["schema"], entity.JsonSchema);
	        this.error = source["error"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace entity {
	
	export class CsvData {
	    fileName: string;
	    location: string;
	    headers: string[];
	
	    static createFrom(source: any = {}) {
	        return new CsvData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fileName = source["fileName"];
	        this.location = source["location"];
	        this.headers = source["headers"];
	    }
	}
	export class JsonSchema {
	    title: string;
	    description: string;
	    properties: {[key: string]: any};
	
	    static createFrom(source: any = {}) {
	        return new JsonSchema(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.description = source["description"];
	        this.properties = source["properties"];
	    }
	}

}

