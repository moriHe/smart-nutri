import {environment as baseEnvironment} from "./environment"
export const environment = {
    production: true,
    ...baseEnvironment,
    frontendBaseUrl: "https://goldfish-app-fid8j.ondigitalocean.app",
    backendBaseUrl: "https://whale-app-lux4j.ondigitalocean.app",
    supabaseUrl: 'https://xmuxozglgkbzuagwjcqt.supabase.co',
    supabaseKey: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InhtdXhvemdsZ2tienVhZ3dqY3F0Iiwicm9sZSI6ImFub24iLCJpYXQiOjE3MDU1ODI1MzEsImV4cCI6MjAyMTE1ODUzMX0.q5ixF3scRGtXG6bZPPM2l8sXYoHrhkBbkVWOHSx1eWs'
}
