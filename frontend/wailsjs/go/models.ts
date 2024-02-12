export namespace invidious {
	
	export class Instance {
	    apiUrl: string;
	    region: string;
	
	    static createFrom(source: any = {}) {
	        return new Instance(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.apiUrl = source["apiUrl"];
	        this.region = source["region"];
	    }
	}
	export class searchAuthorThumbnail {
	    url: string;
	    width: number;
	    height: number;
	
	    static createFrom(source: any = {}) {
	        return new searchAuthorThumbnail(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.width = source["width"];
	        this.height = source["height"];
	    }
	}
	export class videoCaption {
	    label: string;
	    languageCode: string;
	    url: string;
	    type: string;
	
	    static createFrom(source: any = {}) {
	        return new videoCaption(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.label = source["label"];
	        this.languageCode = source["languageCode"];
	        this.url = source["url"];
	        this.type = source["type"];
	    }
	}
	export class videoFormatStream {
	    url: string;
	    iTag: string;
	    type: string;
	    quality: string;
	    fps: number;
	    container: string;
	    encoding: string;
	    qualityLabel: string;
	    resultion: string;
	    size: string;
	
	    static createFrom(source: any = {}) {
	        return new videoFormatStream(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.iTag = source["iTag"];
	        this.type = source["type"];
	        this.quality = source["quality"];
	        this.fps = source["fps"];
	        this.container = source["container"];
	        this.encoding = source["encoding"];
	        this.qualityLabel = source["qualityLabel"];
	        this.resultion = source["resultion"];
	        this.size = source["size"];
	    }
	}
	export class videoAdaptiveFormat {
	    index: string;
	    bitrate: string;
	    init: string;
	    url: string;
	    iTag: string;
	    type: string;
	    clen: string;
	    lmt: string;
	    projectionType: string;
	    fps: number;
	    container: string;
	    encoding: string;
	    qualityLabel: string;
	    resolution: string;
	    audioQuality: string;
	    audioSampleRate: number;
	    audioChannels: number;
	
	    static createFrom(source: any = {}) {
	        return new videoAdaptiveFormat(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.index = source["index"];
	        this.bitrate = source["bitrate"];
	        this.init = source["init"];
	        this.url = source["url"];
	        this.iTag = source["iTag"];
	        this.type = source["type"];
	        this.clen = source["clen"];
	        this.lmt = source["lmt"];
	        this.projectionType = source["projectionType"];
	        this.fps = source["fps"];
	        this.container = source["container"];
	        this.encoding = source["encoding"];
	        this.qualityLabel = source["qualityLabel"];
	        this.resolution = source["resolution"];
	        this.audioQuality = source["audioQuality"];
	        this.audioSampleRate = source["audioSampleRate"];
	        this.audioChannels = source["audioChannels"];
	    }
	}
	export class videoAuthorThumbnail {
	    url: string;
	    width: number;
	    height: number;
	
	    static createFrom(source: any = {}) {
	        return new videoAuthorThumbnail(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.width = source["width"];
	        this.height = source["height"];
	    }
	}
	export class videoVideoThumbnail {
	    quality: string;
	    url: string;
	    width: number;
	    height: number;
	
	    static createFrom(source: any = {}) {
	        return new videoVideoThumbnail(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.quality = source["quality"];
	        this.url = source["url"];
	        this.width = source["width"];
	        this.height = source["height"];
	    }
	}
	export class Video {
	    title: string;
	    videoId: string;
	    videoThumbnails: videoVideoThumbnail[];
	    description: string;
	    descriptionHtml: string;
	    published: number;
	    publishedText: string;
	    keywords: string[];
	    viewCount: number;
	    likeCount: number;
	    dislikeCount: number;
	    paid: boolean;
	    premium: boolean;
	    isFamilyFriendly: boolean;
	    allowedRegions: string[];
	    genre: string;
	    genreUrl: string;
	    author: string;
	    authorId: string;
	    authorUrl: string;
	    authorThumbnails: videoAuthorThumbnail[];
	    subCountText: string;
	    lengthSeconds: number;
	    allowRatings: boolean;
	    rating: number;
	    isListed: boolean;
	    liveNow: boolean;
	    isUpcoming: boolean;
	    premiereTimestamp: number;
	    dashUrl: string;
	    adaptiveFormats: videoAdaptiveFormat[];
	    formatStreams: videoFormatStream[];
	    captions: videoCaption[];
	    recommendedVideos: Video[];
	    tubedVideoThumbnailUrl: string;
	    tubedAuthorThumbnailUrl: string;
	
	    static createFrom(source: any = {}) {
	        return new Video(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.videoId = source["videoId"];
	        this.videoThumbnails = this.convertValues(source["videoThumbnails"], videoVideoThumbnail);
	        this.description = source["description"];
	        this.descriptionHtml = source["descriptionHtml"];
	        this.published = source["published"];
	        this.publishedText = source["publishedText"];
	        this.keywords = source["keywords"];
	        this.viewCount = source["viewCount"];
	        this.likeCount = source["likeCount"];
	        this.dislikeCount = source["dislikeCount"];
	        this.paid = source["paid"];
	        this.premium = source["premium"];
	        this.isFamilyFriendly = source["isFamilyFriendly"];
	        this.allowedRegions = source["allowedRegions"];
	        this.genre = source["genre"];
	        this.genreUrl = source["genreUrl"];
	        this.author = source["author"];
	        this.authorId = source["authorId"];
	        this.authorUrl = source["authorUrl"];
	        this.authorThumbnails = this.convertValues(source["authorThumbnails"], videoAuthorThumbnail);
	        this.subCountText = source["subCountText"];
	        this.lengthSeconds = source["lengthSeconds"];
	        this.allowRatings = source["allowRatings"];
	        this.rating = source["rating"];
	        this.isListed = source["isListed"];
	        this.liveNow = source["liveNow"];
	        this.isUpcoming = source["isUpcoming"];
	        this.premiereTimestamp = source["premiereTimestamp"];
	        this.dashUrl = source["dashUrl"];
	        this.adaptiveFormats = this.convertValues(source["adaptiveFormats"], videoAdaptiveFormat);
	        this.formatStreams = this.convertValues(source["formatStreams"], videoFormatStream);
	        this.captions = this.convertValues(source["captions"], videoCaption);
	        this.recommendedVideos = this.convertValues(source["recommendedVideos"], Video);
	        this.tubedVideoThumbnailUrl = source["tubedVideoThumbnailUrl"];
	        this.tubedAuthorThumbnailUrl = source["tubedAuthorThumbnailUrl"];
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
	export class searchVideoThumbnail {
	    quality: string;
	    url: string;
	    width: number;
	    height: number;
	
	    static createFrom(source: any = {}) {
	        return new searchVideoThumbnail(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.quality = source["quality"];
	        this.url = source["url"];
	        this.width = source["width"];
	        this.height = source["height"];
	    }
	}
	export class SearchItem {
	    type: string;
	    title: string;
	    videoId: string;
	    author: string;
	    authorId: string;
	    authorUrl: string;
	    viewCount: number;
	    videoThumbnails: searchVideoThumbnail[];
	    description: string;
	    descriptionHtml: string;
	    published: number;
	    publishedText: string;
	    lengthSeconds: number;
	    liveNow: boolean;
	    paid: boolean;
	    premium: boolean;
	    authorVerified: boolean;
	    playlistId: string;
	    playlistThumbnail: string;
	    videos: Video[];
	    authorThumbnails: searchAuthorThumbnail[];
	    autoGenerated: boolean;
	    subCount: number;
	    videoCount: number;
	
	    static createFrom(source: any = {}) {
	        return new SearchItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.title = source["title"];
	        this.videoId = source["videoId"];
	        this.author = source["author"];
	        this.authorId = source["authorId"];
	        this.authorUrl = source["authorUrl"];
	        this.viewCount = source["viewCount"];
	        this.videoThumbnails = this.convertValues(source["videoThumbnails"], searchVideoThumbnail);
	        this.description = source["description"];
	        this.descriptionHtml = source["descriptionHtml"];
	        this.published = source["published"];
	        this.publishedText = source["publishedText"];
	        this.lengthSeconds = source["lengthSeconds"];
	        this.liveNow = source["liveNow"];
	        this.paid = source["paid"];
	        this.premium = source["premium"];
	        this.authorVerified = source["authorVerified"];
	        this.playlistId = source["playlistId"];
	        this.playlistThumbnail = source["playlistThumbnail"];
	        this.videos = this.convertValues(source["videos"], Video);
	        this.authorThumbnails = this.convertValues(source["authorThumbnails"], searchAuthorThumbnail);
	        this.autoGenerated = source["autoGenerated"];
	        this.subCount = source["subCount"];
	        this.videoCount = source["videoCount"];
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
	export class SearchOption {
	    page: number;
	    sortBy: string;
	    date: string;
	    duration: string;
	    type: string;
	    features: string[];
	    region: string;
	
	    static createFrom(source: any = {}) {
	        return new SearchOption(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.page = source["page"];
	        this.sortBy = source["sortBy"];
	        this.date = source["date"];
	        this.duration = source["duration"];
	        this.type = source["type"];
	        this.features = source["features"];
	        this.region = source["region"];
	    }
	}
	export class TrendingOption {
	    type: string;
	    region: string;
	
	    static createFrom(source: any = {}) {
	        return new TrendingOption(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.region = source["region"];
	    }
	}
	
	
	
	
	
	
	

}

