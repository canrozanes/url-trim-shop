import { styled } from "@material-ui/core/styles"
import TextField from "@material-ui/core/TextField"
import Link from "@material-ui/core/Link"
import Button from "@material-ui/core/Button"

export const StyledButton = styled(Button)(({ theme }) => ({
  padding: "16px 0",
  margin: "16px 0",
  [theme.breakpoints.up("md")]: {
    marginLeft: "20px",
    width: "calc(100% - 20px)",
    margin: "0",
  },
}))

export const SmallButton = styled(Button)(({ theme }) => ({
  width: "100px",
  alignSelf: "center",
  [theme.breakpoints.up("md")]: {
    marginLeft: "20px",
  },
}))

export const UrlsContainer = styled("div")(({ theme }) => ({
  borderRadius: "4px",
  border: `2px solid ${theme.palette.primary.main}`,
  padding: "16px 20px",
  margin: "16px 0",
  backgroundColor: theme.palette.primary.light,
  [theme.breakpoints.up("md")]: {
    display: "flex",
    justifyContent: "space-between",
    alignItems: "center",
  },
}))

export const UrlCard = styled("div")(({ theme }) => ({
  display: "flex",
  flexDirection: "column",
  marginBottom: "16px",
  [theme.breakpoints.up("md")]: {
    flexDirection: "row",
    alignItems: "center",
    marginBottom: "0",
  },
}))

export const StyledTextField = styled(TextField)({
  maxWidth: "100%",
})

export const StyledAbsoluteLink = styled("a")(({ theme }) => ({
  [theme.breakpoints.up("md")]: {
    marginLeft: "10px",
  },
}))

export const StyledAlertContainer = styled("div")({
  margin: "16px 0",
})
