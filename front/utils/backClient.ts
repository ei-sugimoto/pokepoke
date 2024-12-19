type BackClient = {
    URL: string;
}

export const newBackClient = (): BackClient => {
    if (process.env.NODE_ENV !== 'production') {
        const url = "http://back:8000";
        return {
            URL: url,
        }
    } else {
        return {
            URL: "",
        }
    }
}