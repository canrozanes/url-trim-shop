import React, { useState, MouseEvent } from "react"
import Grid from "@material-ui/core/Grid"
import Typography from "@material-ui/core/Typography"
import { Alert } from "@material-ui/lab"
import CircularProgress from "@material-ui/core/CircularProgress"
import postData from "../../services/index"
import {
  StyledButton,
  SmallButton,
  UrlCard,
  UrlsContainer,
  StyledTextField,
  StyledAbsoluteLink,
  StyledAlertContainer,
} from "./styles"

const isEmpty = (string: string) =>
  typeof string === "string" && string.length === 0

const validUrlRegex = /(?:https?:\/\/)?(?:[a-zA-Z0-9.-]+?\.(?:[a-zA-Z])|\d+\.\d+\.\d+\.\d+)/

export const isValidUrl = (string: string): boolean => {
  return string.match(validUrlRegex) !== null
}

const Form = () => {
  const [urlField, setUrlField] = useState("")
  const [originalUrlField, setOriginalUrlField] = useState("")
  const [shortenedUrl, setShortenedUrl] = useState("")
  const [isFetching, setIsFetching] = useState(false)
  const [isClicked, setIsClicked] = useState(false)
  const [isUrlError, setIsUrlError] = useState(false)

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setUrlField(event.target.value)
  }
  const handleClick = async (e: MouseEvent<HTMLButtonElement>) => {
    e.preventDefault()
    setIsUrlError(false)

    if (!isValidUrl(urlField)) {
      setIsUrlError(true)
      return
    }
    setIsFetching(true)
    try {
      const res = await postData(urlField)
      setShortenedUrl(`${window.location.hostname}/${res.body.Hash}`)
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
      {isUrlError && (
        <StyledAlertContainer>
          <Alert severity="error">Please submit a valid url</Alert>
        </StyledAlertContainer>
      )}
      {!isEmpty(shortenedUrl) && (
        <UrlsContainer>
          <UrlCard>
            <Typography variant="body1" component="span">
              Original URL:
            </Typography>{" "}
            <Typography variant="body1" component="span">
              <StyledAbsoluteLink color="textPrimary">{`http://${originalUrlField}`}</StyledAbsoluteLink>
            </Typography>
          </UrlCard>
          <UrlCard>
            <Typography variant="body1" component="span">
              {"Shortened URL:  "}
            </Typography>
            <Typography variant="body1" component="span">
              <StyledAbsoluteLink
                href={`http://${shortenedUrl}`}
                target="_blank"
              >
                {shortenedUrl}
              </StyledAbsoluteLink>
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
