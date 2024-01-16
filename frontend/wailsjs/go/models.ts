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
	export class HeaderDescriptor {
	    header: string;
	    schemaIdx: number;
	    fieldIdx: number;
	
	    static createFrom(source: any = {}) {
	        return new HeaderDescriptor(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.header = source["header"];
	        this.schemaIdx = source["schemaIdx"];
	        this.fieldIdx = source["fieldIdx"];
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

