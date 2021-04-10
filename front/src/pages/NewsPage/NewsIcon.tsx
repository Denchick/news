import React from "react";
import { Image } from "react-bootstrap";

interface INewsIconProps {
    src: string;
}

const NewsIcon = ({ src }: INewsIconProps) => <Image src={src} rounded width={32} height={32} />;

export default NewsIcon;