export interface IRequester {
  url: string;
  method: Http_Methods;
  data?: unknown;
  headers?: { [key: string]: string };
}

export type Http_Methods =
  | "GET"
  | "POST"
  | "PUT"
  | "DELETE"
  | "OPTION"
  | "PATCH";
