import { cn } from "@/utils";
import { Link as ReactRouterLink } from "@tanstack/react-router";
import type { ComponentPropsWithoutRef } from "react";

export const Link = (
  props: ComponentPropsWithoutRef<typeof ReactRouterLink>,
) => {
  return (
    <ReactRouterLink
      {...props}
      className={cn(
        "hover:underline",
        "text-blue-500",
        "underline-offset-2",
        "decoration-blue-500",
        ...(props.className ?? ""),
      )}
    >
      {props.children}
    </ReactRouterLink>
  );
};
