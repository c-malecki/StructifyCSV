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
	
	export class CsvHeader {
	    column: string;
	    header: string;
	
	    static createFrom(source: any = {}) {
	        return new CsvHeader(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.column = source["column"];
	        this.header = source["header"];
	    }
	}
	export class CsvFileData {
	    fileName: string;
	    location: string;
	    headers: CsvHeader[];
	    referenceRows: string[][];
	
	    static createFrom(source: any = {}) {
	        return new CsvFileData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fileName = source["fileName"];
	        this.location = source["location"];
	        this.headers = this.convertValues(source["headers"], CsvHeader);
	        this.referenceRows = source["referenceRows"];
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
	
	export class RowError {
	    row: number;
	    column: string;
	    propertyKey: string;
	    propertyType: string;
	    value: any;
	    errorType: string;
	    errorMessage: string;
	
	    static createFrom(source: any = {}) {
	        return new RowError(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.row = source["row"];
	        this.column = source["column"];
	        this.propertyKey = source["propertyKey"];
	        this.propertyType = source["propertyType"];
	        this.value = source["value"];
	        this.errorType = source["errorType"];
	        this.errorMessage = source["errorMessage"];
	    }
	}
	export class CsvProcessingReport {
	    successes: number[];
	    rowErrors: RowError[];
	
	    static createFrom(source: any = {}) {
	        return new CsvProcessingReport(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.successes = source["successes"];
	        this.rowErrors = this.convertValues(source["rowErrors"], RowError);
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
	export class PropertySchema {
	    type: string;
	    oneOf: PropertySchema[];
	    minProperties?: number;
	    maxProperties?: number;
	    required: string[];
	    properties: {[key: string]: PropertySchema};
	    minItems?: number;
	    maxItems?: number;
	    items: any;
	    minLength?: number;
	    maxLength?: number;
	    numMinimum?: number;
	    numMaximum?: number;
	    intMinimum?: number;
	    intMaximum?: number;
	    csvHeaderIndex: any;
	
	    static createFrom(source: any = {}) {
	        return new PropertySchema(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.oneOf = this.convertValues(source["oneOf"], PropertySchema);
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
	    properties: {[key: string]: PropertySchema};
	
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

