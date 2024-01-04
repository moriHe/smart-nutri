export interface Response<T> {
    data: T
    // todo return status from go. Not working right now
    status: number
}