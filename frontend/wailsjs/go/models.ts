export namespace main {
	
	export class Caption {
	    label: string;
	    language: string;
	    url: string;
	
	    static createFrom(source: any = {}) {
	        return new Caption(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.label = source["label"];
	        this.language = source["language"];
	        this.url = source["url"];
	    }
	}
	export class Video {
	    title: string;
	    id: string;
	    url: string;
	    embedUrl: string;
	    thumbnailUrl: string;
	    author: string;
	    authorId: string;
	    authorUrl: string;
	    authorThumbnailUrl: string;
	    description: string;
	    descriptionHtml: string;
	    published: number;
	    publishedText: string;
	    genre: string;
	    liveNow: boolean;
	    subCountText: string;
	    lengthSeconds: number;
	    allowRatings: boolean;
	    rating: number;
	    isListed: boolean;
	    isUpcoming: boolean;
	    viewCount: number;
	    viewCountText: string;
	    likeCount: number;
	    dislikeCount: number;
	    paid: boolean;
	    premium: boolean;
	    isFamilyFriendly: boolean;
	    captions: Caption[];
	    recommendedVideos: Video[];
	
	    static createFrom(source: any = {}) {
	        return new Video(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.id = source["id"];
	        this.url = source["url"];
	        this.embedUrl = source["embedUrl"];
	        this.thumbnailUrl = source["thumbnailUrl"];
	        this.author = source["author"];
	        this.authorId = source["authorId"];
	        this.authorUrl = source["authorUrl"];
	        this.authorThumbnailUrl = source["authorThumbnailUrl"];
	        this.description = source["description"];
	        this.descriptionHtml = source["descriptionHtml"];
	        this.published = source["published"];
	        this.publishedText = source["publishedText"];
	        this.genre = source["genre"];
	        this.liveNow = source["liveNow"];
	        this.subCountText = source["subCountText"];
	        this.lengthSeconds = source["lengthSeconds"];
	        this.allowRatings = source["allowRatings"];
	        this.rating = source["rating"];
	        this.isListed = source["isListed"];
	        this.isUpcoming = source["isUpcoming"];
	        this.viewCount = source["viewCount"];
	        this.viewCountText = source["viewCountText"];
	        this.likeCount = source["likeCount"];
	        this.dislikeCount = source["dislikeCount"];
	        this.paid = source["paid"];
	        this.premium = source["premium"];
	        this.isFamilyFriendly = source["isFamilyFriendly"];
	        this.captions = this.convertValues(source["captions"], Caption);
	        this.recommendedVideos = this.convertValues(source["recommendedVideos"], Video);
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

