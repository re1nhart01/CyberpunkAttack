import { defaultTo } from "ramda";

export function clearAfterMs(cb: (v: string) => void, v: string, ms = 1500) {
    cb(v);
    setTimeout(() => {
      cb("");
    }, ms);
  }

  export function sleep(ms: number) {
      return new Promise((resolve) => setTimeout(resolve, ms));
    }

    export function unwrapInto<T>(object: T, v: { [key: string]: unknown }) {
        return { ...object, ...v };
      }



      export function arrayToDictionary<T>(
          arr: Array<T>,
          k: keyof T
        ): { [key: string]: T } {
          const result: { [key: string]: T } = {};
          for (const i of arr) {
            const keyOfObject = i[k];
            result[keyOfObject as string] = i;
          }

          return result;
        }



        export const isPrimitive = <T>(data: T) =>
          ["string", "boolean", "number", "symbol", "undefined", "null"].includes(
            typeof data
          );



        export function fulfilledOr<T>(data: PromiseSettledResult<T>, or: T) {
          return data.status === "fulfilled" ? data.value : or;
        }

        export const valueOrOther = (key: string, def: string) => (props: object) =>
          defaultTo(def, (<never>props)[key]);

        export const valueOrZero =
          (key: string, def = 0) =>
          (props: object) =>
            defaultTo(def, (<never>props)[key]);