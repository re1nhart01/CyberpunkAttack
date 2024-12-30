import { TypographyComponents } from "../../typography/typography.styles";
import React, { useState, useEffect } from "react";
import styled from "styled-components";

const OuterContainer = styled.div`
  margin-top: 30px;
  margin-bottom: 30px;
`;
// Styled Components
const UploadContainer = styled.div`
  width: 100%;
  padding: 20px;
  border: 2px dashed #00f7ff;
  border-radius: 10px;
  text-align: center;
  cursor: pointer;
  transition: 0.3s;
  background: linear-gradient(145deg, #0f172a, #1e293b);
  box-shadow: 0 4px 6px rgba(0, 247, 255, 0.2), 0 1px 3px rgba(0, 247, 255, 0.4);

  &:hover {
    box-shadow: 0 0 10px #00f7ff, 0 0 20px #00f7ff, 0 0 30px #00f7ff;
    border-color: #00f7ff;
  }

  &.dragover {
    background: linear-gradient(145deg, #1e293b, #0f172a);
    box-shadow: 0 0 15px #ff00ff, 0 0 30px #ff00ff;
    border-color: #ff00ff;
  }
`;

const Input = styled.input`
  display: none;
`;

const PreviewWrapper = styled.div`
  position: relative;
  margin-top: 15px;
  width: 100%;
  max-height: 200px;
`;

const PreviewImage = styled.img`
  width: 100%;
  max-height: 200px;
  object-fit: contain;
  border: 1px solid #00f7ff;
  border-radius: 5px;
  box-shadow: 0 0 10px #00f7ff, 0 0 20px #00f7ff;
`;

const RemoveButton = styled.div`
  position: absolute;
  top: 0;
  right: 0;
  padding: 5px 10px;
  background: rgba(0, 0, 0, 0.8);
  color: #ff3860;
  font-size: 12px;
  cursor: pointer;
  opacity: 0;
  transition: 0.3s;
  border-radius: 0 5px 0 5px;

  ${PreviewWrapper}:hover & {
    opacity: 1;
  }

  &:hover {
    color: #ff00ff;
  }
`;

const Status = styled.p<{ error: boolean }>`
  margin-top: 10px;
  font-size: 14px;
  color: ${({ error }) => (error ? "#ff3860" : "#00f7ff")};
  text-shadow: 0 0 5px ${({ error }) => (error ? "#ff3860" : "#00f7ff")};
  letter-spacing: 1px;
  text-transform: uppercase;
`;

const { Text14Zekton700, Text16Zekton400, Text11Zekton400 } =
  TypographyComponents;

// Props
interface FileUploaderProps {
  value: File | null;
  setValue: (file: File | null) => void;
}

const DragNDropInput: React.FC<FileUploaderProps> = ({ value, setValue }) => {
  const [dragging, setDragging] = useState<boolean>(false);
  const [imagePreview, setImagePreview] = useState<string>("");
  const [statusMessage, setStatusMessage] = useState<string>("");
  const [error, setError] = useState<boolean>(false);

  // Generate preview when file changes
  useEffect(() => {
    if (value) {
      const reader = new FileReader();
      reader.onload = (e) => {
        setImagePreview(e.target?.result as string);
      };
      reader.readAsDataURL(value);
    } else {
      setImagePreview("");
    }
  }, [value]);

  // File validation and setting
  const validateAndSetFile = (file: File | null) => {
    if (file && file.type.startsWith("image/")) {
      setValue(file);
      setError(false);
    } else {
      setStatusMessage("Invalid file. Please upload an image.");
      setError(true);
      setValue(null);
    }
  };

  // Remove image
  const removeImage = () => {
    setValue(null);
    setImagePreview("");
    setError(false);
  };

  // Drag-and-Drop Handlers
  const handleDrop = (event: React.DragEvent<HTMLDivElement>) => {
    event.preventDefault();
    setDragging(false);
    const file = event.dataTransfer.files[0] || null;
    validateAndSetFile(file);
  };

  const handleDragOver = (event: React.DragEvent<HTMLDivElement>) => {
    event.preventDefault();
    setDragging(true);
  };

  const handleDragLeave = () => {
    setDragging(false);
  };

  const handleFileSelect = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0] || null;
    validateAndSetFile(file);
  };

  return (
    <OuterContainer>
      <UploadContainer
        className={dragging ? "dragover" : ""}
        onDragOver={handleDragOver}
        onDragLeave={handleDragLeave}
        onDrop={handleDrop}
        onClick={() => document.getElementById("fileInput")?.click()}
      >
        <Text16Zekton400>
          Drop your photo here or click to upload
        </Text16Zekton400>
        <br />
        <Text11Zekton400>
          (only if you pack include custom paint)
        </Text11Zekton400>
      </UploadContainer>

      <Input
        id="fileInput"
        type="file"
        accept="image/*"
        onChange={handleFileSelect}
      />

      {imagePreview && (
        <PreviewWrapper>
          <PreviewImage src={imagePreview} alt="Preview" />
          <RemoveButton onClick={removeImage}>
            <Text14Zekton700>Remove</Text14Zekton700>
          </RemoveButton>
        </PreviewWrapper>
      )}
      {statusMessage && (
        <Status error={error}>
          <Text16Zekton400>{statusMessage}</Text16Zekton400>
        </Status>
      )}
    </OuterContainer>
  );
};

export default DragNDropInput;
