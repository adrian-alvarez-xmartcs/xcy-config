import { authenticated_get, authenticated_post, basic_get, basic_post } from "../utils/fetch/fetch";

// const config = `${document.location.protocol}//${document.location.host}`;

const BASE = "http://localhost:8080";

interface User {
    username: string;
    password: string;
}

interface Token {
    guid: string;
    token: string;
}

interface TemplateProps {
    [key: string]: any;
}

interface CheckInTemplate {
    id: string;
    name: string;
    description: string;
    templateInfo: {
        props: TemplateProps;
    };
}

interface CheckInInstance {
    instanceId: string;
    name: string;
    description: string;
    templateId: string;
}

const ApiBackend = {
    Auth: {
        Login: async (user: User) => {
            const url = `${BASE}/api/auth/login`;
            return await basic_post(url, user);
        },
        Logout: async (guid: string, token: string) => {
            const url = `${BASE}/api/auth/logout`;
            const body: Token = { guid, token };
            return await authenticated_post(url, body);
        },
        // CheckTokenValid: async () => {
        //     const url = `${BASE}/api/auth/token`;
        //     return await authenticated_post(url);
        // }
    },
    Template: {
        GetDetailedById: async (id: string) => {
            const url = `${BASE}/api/templates/get?id=${id}`;
            return await authenticated_get(url);
        },
        Browse: async () => {
            const url = `${BASE}/api/templates/browse`;
            return await authenticated_get(url);
        },
        CheckIn: async (id: string, name: string, description: string, props: TemplateProps) => {
            const url = `${BASE}/api/templates/check-in`;
            const body: CheckInTemplate = {
                id,
                name,
                description,
                templateInfo: { props }
            };
            return await authenticated_post(url, body);
        },
        CheckOut: async (idTemplate: string) => {
            const url = `${BASE}/api/templates/check-out`;
            const body = { id: idTemplate };
            return await authenticated_post(url, body);
        },
        Discard: async (idTemplate: string) => {
            const url = `${BASE}/api/templates/discard`;
            const body = { id: idTemplate };
            return await authenticated_post(url, body);
        },
        Delete: async (idTemplate: string) => {
            const url = `${BASE}/api/templates/delete`;
            const body = { id: idTemplate };
            return await authenticated_post(url, body);
        },
        Revision: async (id: string) => {
            const url = `${BASE}/api/templates/revisions?id=${id}`;
            return await authenticated_get(url);
        }
    },
    Instance: {
        GetChildrenById: async (id: string) => {
            const url = `${BASE}/api/instances/browse?parentId=${id}`;
            return await authenticated_get(url);
        },
        GetDetailedById: async (id: string) => {
            const url = `${BASE}/api/instances/get?id=${id}`;
            return await authenticated_get(url);
        },
        Find: async (path: string) => {
            const url = `${BASE}/api/instances/find?path=${path}`;
            return await authenticated_get(url);
        },
        Delete: async (idInstance: string) => {
            const url = `${BASE}/api/instances/delete`;
            const body = { id: idInstance };
            return await authenticated_post(url, body);
        },
        Move: async (instanceId: string, parentId: string) => {
            const url = `${BASE}/api/instances/move`;
            const body = { instanceId, parentId };
            return await authenticated_post(url, body);
        },
        CheckOut: async (idInstance: string) => {
            const url = `${BASE}/api/instances/check-out`;
            const body = { id: idInstance };
            return await authenticated_post(url, body);
        },
        CheckIn: async (instanceId: string, name: string, description: string, templateId: string) => {
            const url = `${BASE}/api/instances/check-in`;
            const body: CheckInInstance = { instanceId, name, description, templateId };
            return await authenticated_post(url, body);
        },
        Discard: async (idInstance: string) => {
            const url = `${BASE}/api/instances/discard`;
            const body = { id: idInstance };
            return await authenticated_post(url, body);
        }
    },
    Workspace: {
        Get: async () => {
            const url = `${BASE}/api/workspace/get`;
            return await basic_get(url);
        },
    }
};

export default ApiBackend;
