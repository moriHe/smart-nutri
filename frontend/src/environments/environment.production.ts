import {environment as baseEnvironment} from "./environment"
export const environment = {
    production: true,
    ...baseEnvironment,
    frontendBaseUrl: "https://goldfish-app-fid8j.ondigitalocean.app",
    backendBaseUrl: "https://whale-app-lux4j.ondigitalocean.app",
}
