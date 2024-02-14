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
	
	export class CsvFileData {
	    fileName: string;
	    location: string;
	    headers: string[];
	
	    static createFrom(source: any = {}) {
	        return new CsvFileData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fileName = source["fileName"];
	        this.location = source["location"];
	        this.headers = source["headers"];
	    }
	}
	export class SchemaProperty {
	    type: string;
	    oneOf: SchemaProperty[];
	    minProperties?: number;
	    maxProperties?: number;
	    required: string[];
	    properties: {[key: string]: SchemaProperty};
	    minItems?: number;
	    maxItems?: number;
	    items?: any;
	    minLength?: number;
	    maxLength?: number;
	    numMinimum?: number;
	    numMaximum?: number;
	    intMinimum?: number;
	    intMaximum?: number;
	    csvHeaderIndex: any;
	
	    static createFrom(source: any = {}) {
	        return new SchemaProperty(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.oneOf = this.convertValues(source["oneOf"], SchemaProperty);
	        this.minProperties = source["minProperties"];
	        this.maxProperties = source["maxProperties"];
	        this.required = source["required"];
	        this.properties = source["properties"];
	        this.minItems = source["minItems"];
	        this.maxItems = source["maxItems"];
	        this.items = source["items"];
	        this.minLength = source["minLength"];
	        this.maxLength = source["maxLength"];
	        this.numMinimum = source["numMinimum"];
	        this.numMaximum = source["numMaximum"];
	        this.intMinimum = source["intMinimum"];
	        this.intMaximum = source["intMaximum"];
	        this.csvHeaderIndex = source["csvHeaderIndex"];
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
	export class JsonSchema {
	    title: string;
	    description: string;
	    type: string;
	    required: string[];
	    properties: {[key: string]: SchemaProperty};
	
	    static createFrom(source: any = {}) {
	        return new JsonSchema(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.description = source["description"];
	        this.type = source["type"];
	        this.required = source["required"];
	        this.properties = source["properties"];
	    }
	}

}

