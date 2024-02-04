import {environment as baseEnvironment} from "./environment"
export const environment = {
    production: false,
    ...baseEnvironment,
    frontendBaseUrl: "http://localhost:4200",
    backendBaseUrl: "http://localhost:8080",
    supabaseUrl: 'https://snedxbbwkahrapdbkqrf.supabase.co',
    supabaseKey: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InNuZWR4YmJ3a2FocmFwZGJrcXJmIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MDcwNjk4MDQsImV4cCI6MjAyMjY0NTgwNH0.vcjI2LBUWfVlxSb8vdnuG4XrwTTXgZIG9Yi43PZPCBk'
};
