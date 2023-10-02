export namespace main {
	
	export class MediaConvertOptions {
	    source: string;
	    target_name: string;
	    target_format: string;
	    target_dir: string;
	    target_video_codec: string;
	    target_audio_codec: string;
	
	    static createFrom(source: any = {}) {
	        return new MediaConvertOptions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.source = source["source"];
	        this.target_name = source["target_name"];
	        this.target_format = source["target_format"];
	        this.target_dir = source["target_dir"];
	        this.target_video_codec = source["target_video_codec"];
	        this.target_audio_codec = source["target_audio_codec"];
	    }
	}

}

