import {environment as baseEnvironment} from "./environment"
export const environment = {
    production: false,
    ...baseEnvironment,
    frontendBaseUrl: "http://localhost:4200",
    backendBaseUrl: "http://localhost:8080"
};
