import type { AxiosRequestHeaders } from "axios";
import axios from "axios";
import { assoc, defaultTo, equals, isNil, pipe } from "ramda";

axios.defaults.baseURL = "";
let counterToLogout = 0;
const maxCountToLogout = 10;

export enum APIErrorsResponse {
  invalid_token = 401,
  inactivate_user = 403,
  technical_works = 540,
  networkError = "Network Error",
  abortingError = "canceled",
}

type AxiosCustomHeaderType = Record<"Authorization", string> &
  Omit<
    Record<"Content-Type", string | RegExpExecArray> &
      Omit<AxiosRequestHeaders, "Content-Type">,
    "Authorization"
  >;

export const newAbortSignal = (timeoutMs: number) => {
  const abortController = new AbortController();
  setTimeout(() => abortController.abort(), timeoutMs || 0);
  return abortController.signal;
};

axios.interceptors.request.use(
  async (config) => {
    const access_token = "";
    if (access_token && isNil(config.headers?.Authorization)) {
      (<AxiosCustomHeaderType>config.headers) = pipe(
        assoc(
          "Content-Type",
          defaultTo("application/json", config.headers?.getContentType?.())
        ),
        assoc("Authorization", `Bearer ${access_token}`)
      )(config.headers);
    }

    return assoc(
      "signal",
      defaultTo(newAbortSignal(15000), config.signal),
      config
    );
  },
  (error) => Promise.reject(error)
);

axios.interceptors.response.use(
  (response) => {
    counterToLogout = 0;
    return response;
  },
  (error) => {
    console.error({ error });

    if (equals(error.message, APIErrorsResponse.abortingError)) {
    }

    if (equals(error.message, APIErrorsResponse.networkError)) {
    }

    if (equals(error?.response?.status, APIErrorsResponse.inactivate_user)) {

    }

    if (equals(error?.response?.status, APIErrorsResponse.technical_works)) {
    }

    if (equals(error?.response?.status, APIErrorsResponse.invalid_token)) {
      counterToLogout++;
      if (counterToLogout >= maxCountToLogout) {

      }
    }

    return Promise.reject(error);
  }
);
