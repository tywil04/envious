export namespace shared {
	
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
	    embedUrl: string;
	    thumbnailUrl: string;
	    author: string;
	    authorId: string;
	    authorAvatarUrl: string;
	    shortDescription: string;
	    description: string;
	    published: number;
	    publishedText: string;
	    genre: string;
	    subCountText: string;
	    lengthSeconds: number;
	    allowRatings: boolean;
	    rating: number;
	    viewCount: number;
	    viewCountText: string;
	    likeCount: number;
	    dislikeCount: number;
	    isListed: boolean;
	    isUpcoming: boolean;
	    isShort: boolean;
	    isPaid: boolean;
	    isPremium: boolean;
	    isFamilyFriendly: boolean;
	    liveNow: boolean;
	    captions: Caption[];
	    recommendedVideos: Video[];
	
	    static createFrom(source: any = {}) {
	        return new Video(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.id = source["id"];
	        this.embedUrl = source["embedUrl"];
	        this.thumbnailUrl = source["thumbnailUrl"];
	        this.author = source["author"];
	        this.authorId = source["authorId"];
	        this.authorAvatarUrl = source["authorAvatarUrl"];
	        this.shortDescription = source["shortDescription"];
	        this.description = source["description"];
	        this.published = source["published"];
	        this.publishedText = source["publishedText"];
	        this.genre = source["genre"];
	        this.subCountText = source["subCountText"];
	        this.lengthSeconds = source["lengthSeconds"];
	        this.allowRatings = source["allowRatings"];
	        this.rating = source["rating"];
	        this.viewCount = source["viewCount"];
	        this.viewCountText = source["viewCountText"];
	        this.likeCount = source["likeCount"];
	        this.dislikeCount = source["dislikeCount"];
	        this.isListed = source["isListed"];
	        this.isUpcoming = source["isUpcoming"];
	        this.isShort = source["isShort"];
	        this.isPaid = source["isPaid"];
	        this.isPremium = source["isPremium"];
	        this.isFamilyFriendly = source["isFamilyFriendly"];
	        this.liveNow = source["liveNow"];
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

