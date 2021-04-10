import CategoryName from "../../data/CategoryName";

export interface IMainPageCategory {
    title: CategoryName;
    path?: string;
}

export const rows: IMainPageCategory[][] = [
    [
        {
            title: CategoryName.News,
            path: "/news"
        },
        {
            title: CategoryName.Development,
            path: "/tech"
        },
        {
            title: CategoryName.PeoplesBlogs,
            path: "/blogs"
        }
    ],
    [
        {
            title: CategoryName.German
        },
        {
            title: CategoryName.Polish
        }
    ],
    [
        {
            title: CategoryName.Frontend
        },
        {
            title: CategoryName.MachineLearning
        },
        {
            title: CategoryName.Backend
        }
    ]
]