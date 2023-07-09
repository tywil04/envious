export namespace main {
	
	export class Video {
	    title: string;
	    videoId: string;
	    thumbnailUrl: string;
	    viewCount: number;
	    author: string;
	    authorUrl: string;
	    published: number;
	    publishedText: string;
	    liveNow: boolean;
	    paid: boolean;
	    premium: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Video(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.videoId = source["videoId"];
	        this.thumbnailUrl = source["thumbnailUrl"];
	        this.viewCount = source["viewCount"];
	        this.author = source["author"];
	        this.authorUrl = source["authorUrl"];
	        this.published = source["published"];
	        this.publishedText = source["publishedText"];
	        this.liveNow = source["liveNow"];
	        this.paid = source["paid"];
	        this.premium = source["premium"];
	    }
	}

}

