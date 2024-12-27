import { FormadjoField } from "../validators/Formadjo";
import { FormadjoFormer } from "../validators/FormadjoFormer";

export interface IShipmentFormTemplate {
  firstName: string;
  patronymic: string;
  lastName: string;
  email: string;
  phoneNumber: string;
  country: string;
  state: string;
  city: string;
  street: string;
  house: string;
  apartment: string;
  zipcode: string;
}

export const shipmentForm = new FormadjoFormer<IShipmentFormTemplate>({
  firstName: new FormadjoField("firstName", "string")
    .setIsRequired(true)
    .setMinLength(2)
    .setMaxLength(50),
  patronymic: new FormadjoField("patronymic", "string")
    .setIsRequired(false)
    .setMaxLength(50),
  lastName: new FormadjoField("lastName", "string")
    .setIsRequired(true)
    .setMinLength(2)
    .setMaxLength(50),
  email: new FormadjoField("email", "string")
    .setIsRequired(true)
    .setRegexpValidation(/^[^\s@]+@[^\s@]+\.[^\s@]+$/),
  phoneNumber: new FormadjoField("phoneNumber", "string")
    .setIsRequired(true)
    .setRegexpValidation(/^\+?[0-9]{7,15}$/),
  country: new FormadjoField("country", "string")
    .setIsRequired(true)
    .setMaxLength(50),
  state: new FormadjoField("state", "string")
    .setIsRequired(true)
    .setMaxLength(50),
  city: new FormadjoField("city", "string")
    .setIsRequired(true)
    .setMaxLength(50),
  street: new FormadjoField("street", "string")
    .setIsRequired(true)
    .setMaxLength(100),
  house: new FormadjoField("house", "string")
    .setIsRequired(true)
    .setMaxLength(10),
  apartment: new FormadjoField("apartment", "string")
    .setIsRequired(false)
    .setMaxLength(10),
  zipcode: new FormadjoField("zipcode", "string")
    .setIsRequired(true)
    .setRegexpValidation(/^[0-9]{5,10}$/),
});
