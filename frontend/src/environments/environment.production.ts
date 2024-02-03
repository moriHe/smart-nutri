import {environment as baseEnvironment} from "./environment"
export const environment = {
    production: true,
    ...baseEnvironment,
    frontendBaseUrl: "https://goldfish-app-fid8j.ondigitalocean.app",
    backendBaseUrl: "http://localhost:8080",
}
