import { FormadjoField } from "./Formadjo";

export type GenericFormadjoValidator<T> = { [key in keyof T]: FormadjoField }

export class FormadjoFormer<T extends object> {
  private readonly __form__: GenericFormadjoValidator<T>;

  constructor(form: GenericFormadjoValidator<T>) {
    this.__form__ = form;
  }

  public get get(): GenericFormadjoValidator<T> {
    return this.__form__;
  }
}
