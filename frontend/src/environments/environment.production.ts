import {environment as baseEnvironment} from "./environment"
export const environment = {
    production: true,
    ...baseEnvironment,
    frontendBaseUrl: "",
    backendBaseUrl: "",
}
