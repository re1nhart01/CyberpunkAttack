import { FormadjoField } from "%/core/validators/Formadjo";
import { FormadjoFormer } from "%/core/validators/FormadjoFormer";

export const userEducationForm = new FormadjoFormer<{
  university: string;
  role: number;
}>({
  university: new FormadjoField("university", "string")
    .setIsRequired(false)
    .setMinLength(2)
    .setMaxLength(40),
  role: new FormadjoField("role", "number").setIsRequired(false),
});
