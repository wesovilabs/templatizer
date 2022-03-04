import axios from "axios";


const API_URL = 'http://localhost:5001/api';

export interface Auth {
    mechanism: string;
    username?: string;
    password?: string;
    token?: string;
}

export interface LoadParametersRequest {
    url: string;
    auth?: Auth;
    branch?: string;
}

export interface Var {
    name: string;
    description: string;
    type: string;
    default: string;
}

export interface ProcessTemplateRequest extends LoadParametersRequest {
    params?: [object];
}

export interface Config {
    version: string;
    mode: string;
    variables: [Var];
}

export const LoadParameters = (request: LoadParametersRequest) => {
    return axios.post(`${API_URL}/parameters`, request)
}

export const ProcessTemplate = (request: ProcessTemplateRequest) => {
    return axios.post(`${API_URL}/template`, request, { responseType: 'blob' })
}
