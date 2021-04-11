export enum CategoryName {
    News = "🌎 News",
    Development =  "👨‍💻 Development",
    PeoplesBlogs = "🙂 People's blogs",
    German = "🇩🇪 German",
    Polish = "🇵🇱 Polish",
    Frontend = "🌈 Front-End",
    MachineLearning = "🤡 Machine Learning",
    Backend = "⚙️ Back-End"
}

export const categoryPaths: {[id in CategoryName]: string | null;} = {
    [CategoryName.News]: "news",
    [CategoryName.Development]: "dev",
    [CategoryName.PeoplesBlogs]: "blogs",
    [CategoryName.German]: "german",
    [CategoryName.Polish]: null,
    [CategoryName.Frontend]: null,
    [CategoryName.MachineLearning]: null,
    [CategoryName.Backend]: null
}