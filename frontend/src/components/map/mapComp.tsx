import React from "react";
import { GoogleMap, LoadScript, Marker } from "@react-google-maps/api";

const defaultMapConfig = {
  gestureHandling: "greedy",
  options: {
    fullscreenControl: false
  },
  mapContainerStyle: {
    height: "36vh",
    width: "100%"
  }
}; 

export default function GMap() {
  return (
    <>
      <LoadScript
        googleMapsApiKey={process.env.NEXT_PUBLIC_GOOGLE_MAPS_API_KEY}
      >
        <GoogleMap {...defaultMapConfig} center={{ lat: 7.494726, lng: 102.612209 }} zoom={7}>
          <Marker position={{ lat: 7.494726, lng: 102.612209 }} draggable={false} />
          <Marker position={{ lat: 8.263407, lng: 102.519507 }} draggable={false} />
          <Marker position={{ lat: 8.147704, lng: 102.255313 }} draggable={false} />

        </GoogleMap>
      </LoadScript>
    </>
  );
}
