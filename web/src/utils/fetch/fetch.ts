import { logErrors } from "../../consts/consts";

interface BasicPostParams {
    [key: string]: any;
}

interface ErrorResponse {
    error: string;
}

const handleError = async (response: Response) => {
    const text = await response.text();
    throw new Error((JSON.parse(text) as ErrorResponse).error);
};

export const basic_get = (url: string) => {
    return fetch(url, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        }
    })
    .then((response) => {
        if (!response.ok) {
            return handleError(response);
        }
        return response.json();
    })
    .catch((error) => {
        if (logErrors) {
            console.error("Error in fetch:", error);
        }
        throw error;
    });
};

export const basic_post = (url: string, params: BasicPostParams) => {
    return fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(params),
    })
    .then((response) => {
        if (!response.ok) {
            return handleError(response);
        }
        return response.json();
    })
    .catch((error) => {
        if (logErrors) {
            console.error("Error in basic_post:", error);
        }
        throw error;
    });
};

export const authenticated_get = (url: string) => {
    return fetch(url, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authentication": localStorage.getItem("auth_token") || "",
        }
    })
    .then((response) => {
        if (response.status !== 200) {
            return handleError(response);
        }
        return response.json();
    })
    .catch((error) => {
        if (logErrors) {
            console.error("Error in fetch:", error);
        }
        throw error;
    });
};

export const authenticated_post = (url: string, params: BasicPostParams) => {
    return fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authentication": localStorage.getItem("auth_token") || "",
        },
        body: JSON.stringify(params)
    })
    .then((response) => {
        if (response.status !== 200) {
            return handleError(response);
        }
        return response.json();
    })
    .catch((error) => {
        if (logErrors) {
            console.error("Error in fetch:", error);
        }
        throw error;
    });
};

