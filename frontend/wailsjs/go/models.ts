export namespace main {
	
	export class ImageConvertOptions {
	    target_format: string;
	
	    static createFrom(source: any = {}) {
	        return new ImageConvertOptions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.target_format = source["target_format"];
	    }
	}
	export class MediaConvertOptions {
	    target_format: string;
	    // Go type: struct { Video string "json:\"video\""; Audio string "json:\"audio\"" }
	    target_bitrate: any;
	
	    static createFrom(source: any = {}) {
	        return new MediaConvertOptions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.target_format = source["target_format"];
	        this.target_bitrate = this.convertValues(source["target_bitrate"], Object);
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

