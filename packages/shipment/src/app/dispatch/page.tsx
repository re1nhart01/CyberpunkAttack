"use client";

import { MainLayout } from "%/components/main-layout/MainLayout";
import Image from "next/image";
import CornerSeparator from "%/assets/section seperator.svg";
import { TextView } from "%/components/text/Text";
import { FormInputView } from "%/components/form-input/CyberInput";

export default function Dispatch() {
  return (
    <div className="app bg-[#ffffff]">
      <TextView
        text="Shipment"
        className="overflow-hidden max-w-[98vw] font-space-armor text-[10vw] font-[400] leading-[15vw] text-left absolute top-[5%] left-[30px] z-[1]"
      />
      <main className="app-container relative z-[2]">
        <Image
          className="cornerSeparator-top"
          src={CornerSeparator}
          alt="cornerSeparator"
        />
        <section className="z-[2] pt-[8%] pl-[60px]">
          <TextView
            text="Personal Info"
            className="overflow-hidden font-orbitron text-black max-w-[98vw] text-[4rem] font-[400] text-left z-[2]"
          />
          <div className="pt-[20px]">
            <FormInputView
              inputProps={{ placeholder: "First Name" }}
              onChange={() => {}}
            />
          </div>
          <div className="pt-[20px]">
            <FormInputView
              inputProps={{ placeholder: "First Name" }}
              onChange={() => {}}
            />
          </div>
        </section>
      </main>
      <Image
        className="cornerSeparator-bottom"
        src={CornerSeparator}
        alt="cornerSeparator"
      />
    </div>
  );
}
