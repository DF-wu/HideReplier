import { AdvancedInfo } from "../types";
import { useQuery } from "@tanstack/react-query";
import { getAdvancedInfo } from "./api";

export function useAdvancedInfo(): AdvancedInfo {
  const { data = {} } = useQuery(["advanced-info"], getAdvancedInfo, {
    refetchOnMount: false,
    refetchOnWindowFocus: false,
  });
  return AdvancedInfo.parse(data);
}
