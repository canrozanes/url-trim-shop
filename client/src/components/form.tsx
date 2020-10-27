import React, { useState, MouseEvent } from "react"
import { styled } from "@material-ui/core/styles"
import TextField from "@material-ui/core/TextField"
import Link from "@material-ui/core/Link"
import Button from "@material-ui/core/Button"
import Grid from "@material-ui/core/Grid"
import Typography from "@material-ui/core/Typography"
import CircularProgress from "@material-ui/core/CircularProgress"
import postData from "../services/index"

const StyledButton = styled(Button)(({ theme }) => ({
  padding: "16px 0",
  margin: "16px 0",
  [theme.breakpoints.up("md")]: {
    marginLeft: "20px",
    width: "calc(100% - 20px)",
    margin: "0",
  },
}))

const SmallButton = styled(Button)(({ theme }) => ({
  width: "100px",
  alignSelf: "center",
  [theme.breakpoints.up("md")]: {
    marginLeft: "20px",
  },
}))

const UrlsContainer = styled("div")(({ theme }) => ({
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

const UrlCard = styled("div")(({ theme }) => ({
  display: "flex",
  flexDirection: "column",
  marginBottom: "16px",
  [theme.breakpoints.up("md")]: {
    flexDirection: "row",
    alignItems: "center",
    marginBottom: "0",
  },
}))

const StyledTextField = styled(TextField)({
  maxWidth: "100%",
})

const StyledLink = styled(Link)(({ theme }) => ({
  [theme.breakpoints.up("md")]: {
    marginLeft: "10px",
  },
}))

const isNotEmpty = (string: string) =>
  typeof string === "string" && string.length > 0

const Form = () => {
  const [urlField, setUrlField] = useState("")
  const [originalUrlField, setOriginalUrlField] = useState("")
  const [shortenedUrl, setShortenedUrl] = useState("")
  const [isFetching, setIsFetching] = useState(false)
  const [isClicked, setIsClicked] = useState(false)

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setUrlField(event.target.value)
  }
  const handleClick = async (e: MouseEvent<HTMLButtonElement>) => {
    setIsFetching(true)
    e.preventDefault()
    try {
      const res = await postData(urlField)
      setShortenedUrl(`${window.location}${res.body.Hash}`)
      setIsFetching(false)
    } catch (e) {
      console.log(e)
      setIsFetching(false)
    }
    setOriginalUrlField(urlField)
  }

  const copyToClipboard = (e: MouseEvent<HTMLButtonElement>) => {
    e.preventDefault()
    setIsClicked(true)
    const tempInput = document.createElement("input")
    tempInput.value = shortenedUrl
    document.body.appendChild(tempInput)
    tempInput.select()
    document.execCommand("copy")
    document.body.removeChild(tempInput)
    setTimeout(() => {
      setIsClicked(false)
    }, 700)
  }

  return (
    <>
      <Grid container component="form" alignItems="center">
        <Grid item xs={12} md={9}>
          <StyledTextField
            label="Shorten your link"
            variant="filled"
            onChange={handleChange}
            fullWidth
          />
        </Grid>
        <Grid item xs={12} md={3}>
          <StyledButton
            type="submit"
            id="submitButton"
            onClick={handleClick}
            variant="contained"
            color="primary"
            disabled={isFetching}
            fullWidth
            endIcon={
              isFetching ? (
                <CircularProgress
                  color="secondary"
                  size={20}
                  aria-describedby="submitButton"
                  aria-busy={true}
                />
              ) : null
            }
          >
            Shorten
          </StyledButton>
        </Grid>
      </Grid>
      {isNotEmpty(shortenedUrl) && (
        <UrlsContainer>
          <UrlCard>
            <Typography variant="body1" component="span">
              Original URL:
            </Typography>{" "}
            <Typography variant="body1" component="span">
              <StyledLink color="textPrimary">{`http://${originalUrlField}`}</StyledLink>
            </Typography>
          </UrlCard>
          <UrlCard>
            <Typography variant="body1" component="span">
              {"Shortened URL:  "}
            </Typography>
            <Typography variant="body1" component="span">
              <StyledLink href={shortenedUrl}>{shortenedUrl}</StyledLink>
            </Typography>
            <SmallButton
              variant="contained"
              color={isClicked ? "default" : "secondary"}
              onClick={copyToClipboard}
            >
              {isClicked ? "Copied!" : "Copy"}
            </SmallButton>
          </UrlCard>
        </UrlsContainer>
      )}
    </>
  )
}

export default Form
