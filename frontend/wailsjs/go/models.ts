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
	export class SchemaProperty {
	    key: string;
	    value: string;
	    path: string;
	
	    static createFrom(source: any = {}) {
	        return new SchemaProperty(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.value = source["value"];
	        this.path = source["path"];
	    }
	}
	export class HeaderDescriptor {
	    isSelected: boolean;
	    headerText: string;
	    headerIndex: number;
	    schemaProperty?: SchemaProperty;
	
	    static createFrom(source: any = {}) {
	        return new HeaderDescriptor(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.isSelected = source["isSelected"];
	        this.headerText = source["headerText"];
	        this.headerIndex = source["headerIndex"];
	        this.schemaProperty = this.convertValues(source["schemaProperty"], SchemaProperty);
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
	export class CsvModel {
	    headerDescriptors: HeaderDescriptor[];
	    usedHeaderIndexes: number[];
	    map: {[key: string]: any};
	
	    static createFrom(source: any = {}) {
	        return new CsvModel(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.headerDescriptors = this.convertValues(source["headerDescriptors"], HeaderDescriptor);
	        this.usedHeaderIndexes = source["usedHeaderIndexes"];
	        this.map = source["map"];
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

