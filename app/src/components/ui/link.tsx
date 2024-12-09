import { cn } from "@/utils";
import { Link as ReactRouterLink } from "@tanstack/react-router";
import type { ComponentPropsWithoutRef } from "react";
import { Button } from "./button";

export const Link = ({
  className,
  ...props
}: ComponentPropsWithoutRef<typeof ReactRouterLink>) => {
  return (
    <Button
      asChild
      variant="link"
      className={cn("px-1", "py-0", "h-auto", className)}
    >
      <ReactRouterLink {...props}>{props.children}</ReactRouterLink>
    </Button>
  );
};
