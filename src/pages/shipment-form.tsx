import type { HeadFC, PageProps } from "gatsby";
import * as React from "react";
import { useCallback, useEffect, useRef, useState } from "react";
import { useTranslation } from "react-i18next";
import { ButtonComponents } from "../components/buttons/Button/components";
import { Html } from "../components/html";
import { ImageViewComponents } from "../components/images/ImageView/styles";
import { InputStyles } from "../components/inputs/styles";
import FooterView from "../components/layout/Footer";
import { FullScreenMenuComponent } from "../components/layout/FullScreenMenu/FullScreenMenu";
import HeaderView from "../components/layout/Header";
import MainLayout from "../components/layout/MainLayout";
import Preloader from "../components/layout/Preloader/Preloader";
import {
  OverrideTypographyComponents,
  TypographyComponents,
} from "../components/typography/typography.styles";
import { service } from "../services";
import {
  IShipmentFormTemplate,
  shipmentForm,
} from "../services/forms/shipment.form";
import {
  FormadjoAsyncSubmitFn,
  FormadjoForm,
} from "../services/validators/FormadjoForm";
import { formValuesType } from "../services/validators/MainFormadjo";
import { HomePageStyles } from "../styles/pageStyles/home.styles";
import { shipmentFormStyles } from "../styles/pageStyles/shipment-form.styles";
import SmallPreloader from "../components/layout/Preloader/SmallPreloader";

const { Container, FormContainer, InputContainer, TextContainer } =
  shipmentFormStyles;

const { PageContainer, Spacer, InnerWrapper, PageContainerForm } = HomePageStyles;
const { Text26Bangers400, Text48Orbitron700, Text24Zekton400 } =
  TypographyComponents;
const { Text16Zekton400Black, Text24Zekton400Link } =
  OverrideTypographyComponents;
const { FormImage1, FormImage2, AlertImage } = ImageViewComponents;
const { PrimaryInput } = InputStyles;
const { ShipmentFormButton } = ButtonComponents;

const workerUrl = "https://small-tooth-64b0.attackcyberpunk.workers.dev";

const safetyLocalStorage = typeof localStorage !== "undefined"
  ? localStorage
  : {
    getItem: (t: string) => {},
    setItem: (t: string, v: string) => {},
    removeItem: (v: string) => {},
  };

const HomePage: React.FC<PageProps> = () => {
  const { t } = useTranslation();
  const params = new URLSearchParams(
    (typeof window !== "undefined"
      ? window
      : { location: { search: "" } }
    ).location.search,
  );
  const isMobile =
    typeof window !== "undefined" ? window.innerWidth < 1024 : false;
  const [isOpen, setIsOpen] = useState(false);
  const isCompleted = safetyLocalStorage?.getItem("isAnswered");
  const [isAnswered, setIsAnswered] = useState(!!isCompleted);
  const [image, setImage] = useState<File | null>(null);
  const [loading, setLoading] = useState(false);

  const onScrollIntoView = (
    arg: "subscribe" | "about" | "trailer" | "start",
  ) => {
    switch (arg) {
      case "start":
      case "subscribe":
      case "about":
      case "trailer":
        window.location.href = "/";
        break;
      default:
    }
  };

  const handleLoadEmailFromParam = (
    func: (k: keyof IShipmentFormTemplate, v: formValuesType) => void,
  ) => {
    try {
      const redirectValue = params.get("invite_id");
      if (redirectValue) {
        const email = decodeURIComponent(escape(atob(redirectValue)));
        email &&
        setTimeout(() => {
          func("email", email);
        }, 0);
      }
    } catch (e) {
      console.log(e);
    }
  };

  const goTo = (url: string) => {
    window.open(url, "_blank");
  };

  const onClean = () => {
    safetyLocalStorage.removeItem("isAnswered");
    setIsAnswered(false);
  };

  const onFormSubmit: FormadjoAsyncSubmitFn<IShipmentFormTemplate> =
    useCallback(async (values: IShipmentFormTemplate) => {
      setLoading(true);
      const formData = new FormData();
      if (image) {
        formData.append("image", image);
      }
      try {
        const responseImgDb = await fetch(`${workerUrl}/imgbb`, {
          method: "POST",
          body: formData,
        });
        const result = await responseImgDb.json();
        await fetch(`${workerUrl}/notion`, {
          method: "POST",
          body: JSON.stringify({
            ...values,
            avatar: result?.data?.display_url ?? "No Avatar Included",
          }),
        });
        setIsAnswered(true);
        localStorage.setItem("isAnswered", "true");
      } catch (e) {
        console.log(e);
      } finally {
        setLoading(false);
      }
    }, [image]);

  useEffect(() => {
    console.log(image);
  }, [image]);

  useEffect(() => {
    service.initServices().then();
  }, []);

  return (
    <MainLayout>
      <HeaderView
        setIsOpen={setIsOpen}
        isOpen={isOpen}
        onScrollIntoView={onScrollIntoView}
      />
      {isMobile && (
        <FullScreenMenuComponent
          setIsOpen={setIsOpen}
          onScrollIntoView={onScrollIntoView}
          isOpen={isOpen}
        />
      )}
      <Preloader />
      <PageContainerForm>
        <Container>
          <FormImage1 />
          {isAnswered ? (
            <FormContainer>
              <Spacer height={120} />
              <AlertImage />
              <InnerWrapper>
                <Text48Orbitron700>
                  {t("shipmentForm.header")}
                </Text48Orbitron700>
              </InnerWrapper>
              <InnerWrapper>
                <Text24Zekton400 color="white">
                  We've received your delivery details. Our team will contact you soon to confirm the shipping.
                  If you have any questions, feel free to reach out at
                  {" "}
                  <Text24Zekton400Link href="mailto:attackcyberpunk@gmail.com">
                    attackcyberpunk@gmail.com
                  </Text24Zekton400Link>
                  .
                </Text24Zekton400>

                <Spacer height={30} />
                <ShipmentFormButton onPress={onClean}>
                  <Text16Zekton400Black>Send Again</Text16Zekton400Black>
                </ShipmentFormButton>
              </InnerWrapper>
            </FormContainer>
          ) : (
            <FormadjoForm<IShipmentFormTemplate>
              removeErrorOnChange
              onLoad={handleLoadEmailFromParam}
              onFinishSubmit={onFormSubmit}
              initialProps={{
                fullName: "",
                email: "",
                phoneNumber: "",
                country: "",
                state: "",
                city: "",
                street: "",
                house: "",
                apartment: "",
                zipcode: "",
              }}
              form={shipmentForm}
            >
              {({ onSubmit, errorsList, updateFormState, values }) => {
                return (
                  <FormContainer>
                    <TextContainer paddingTop={0}>
                      <Text26Bangers400 color="white">
                        User Information
                      </Text26Bangers400>
                    </TextContainer>
                    <InputContainer>
                      <PrimaryInput
                        type="text"
                        disabled={!!params.get("invite_id")}
                        maxLength={150}
                        value={values.email}
                        onChange={({ target }) =>
                          updateFormState("email", target.value)}
                        placeholder="Email*"
                        $isError={errorsList.email.isError}
                      />
                    </InputContainer>
                    <InputContainer>
                      <PrimaryInput
                        type="text"
                        maxLength={150}
                        value={values.phoneNumber}
                        onChange={({ target }) =>
                          updateFormState("phoneNumber", target.value)}
                        placeholder="Phone number*"
                        $isError={errorsList.phoneNumber.isError}
                      />
                    </InputContainer>
                    <InputContainer>
                      <PrimaryInput
                        type="text"
                        maxLength={150}
                        value={values.fullName}
                        onChange={({ target }) =>
                          updateFormState("fullName", target.value)}
                        placeholder="Full Name*"
                        $isError={errorsList.fullName.isError}
                      />
                    </InputContainer>
                    <TextContainer paddingTop={32}>
                      <Text26Bangers400 color="white">
                        Shipping Address
                      </Text26Bangers400>
                    </TextContainer>
                    <InputContainer>
                      <PrimaryInput
                        type="text"
                        maxLength={150}
                        value={values.country}
                        onChange={({ target }) =>
                          updateFormState("country", target.value)}
                        placeholder="Country*"
                        $isError={errorsList.country.isError}
                      />
                    </InputContainer>
                    <InputContainer>
                      <PrimaryInput
                        type="text"
                        maxLength={150}
                        value={values.state}
                        onChange={({ target }) =>
                          updateFormState("state", target.value)}
                        placeholder="State/Region*"
                        $isError={errorsList.state.isError}
                      />
                      <PrimaryInput
                        type="text"
                        maxLength={150}
                        value={values.city}
                        onChange={({ target }) =>
                          updateFormState("city", target.value)}
                        placeholder="City*"
                        $isError={errorsList.city.isError}
                      />
                    </InputContainer>
                    <InputContainer>
                      <PrimaryInput
                        type="text"
                        maxLength={150}
                        value={values.street}
                        onChange={({ target }) =>
                          updateFormState("street", target.value)}
                        placeholder="Street*"
                        $isError={errorsList.street.isError}
                      />
                      <PrimaryInput
                        type="text"
                        maxLength={150}
                        value={values.house}
                        onChange={({ target }) =>
                          updateFormState("house", target.value)}
                        placeholder="House"
                        $isError={errorsList.house.isError}
                      />
                      <PrimaryInput
                        type="text"
                        maxLength={150}
                        value={values.apartment}
                        onChange={({ target }) =>
                          updateFormState("apartment", target.value)}
                        placeholder="Apartment"
                        $isError={errorsList.apartment.isError}
                      />
                      <PrimaryInput
                        type="text"
                        maxLength={150}
                        value={values.zipcode}
                        onChange={({ target }) =>
                          updateFormState("zipcode", target.value)}
                        placeholder="Zip Code*"
                        $isError={errorsList.zipcode.isError}
                      />
                    </InputContainer>
                    {/* <DragNDropInput value={image} setValue={setImage} /> */}
                    <ShipmentFormButton inert={loading} onPress={onSubmit}>
                      {loading ? <SmallPreloader /> : <Text16Zekton400Black>{t("submit")}</Text16Zekton400Black>}
                    </ShipmentFormButton>
                  </FormContainer>
                );
              }}
            </FormadjoForm>
          )}
          <FormImage2 />
        </Container>
        <Spacer height={100} />
        <FooterView />
      </PageContainerForm>
    </MainLayout>
  );
};

export default HomePage;

export const Head: HeadFC = () => (
  <Html
    meta="This board game is a cooperative team techno fight game where 2 - 8 players can clash in a battle. Join the battle as a resistance hacker or take control of the corporation as its boss."
    title="Cyberpunk Attack - cooperative team techno fight game"
  />
);
