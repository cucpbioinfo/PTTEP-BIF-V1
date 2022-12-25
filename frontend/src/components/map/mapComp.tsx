import { GoogleMap, LoadScript, Marker } from "@react-google-maps/api";
import React, { useEffect, useState } from 'react'
import { useRouter } from 'next/router'
import { listLocation } from "api/dashboard/listLocation";
import { listCenter } from "api/dashboard/listCenter";

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
  const router = useRouter()
  let tmplat:any
  let tmplong:any
  let tmpcenlat:any
  let tmpcenlong:any
  const [location, setLocation] = useState([])
  const [center, setCenter] = useState([])
  const fetchLocation = async (assetId?: string) => {
    const { data } = await listLocation(assetId)
    setLocation(data)
  }
  const fetchCenter = async (assetId?: string) => {
    const { data } = await listCenter(assetId)
    setCenter(data)
  }
  useEffect(() => {
    fetchLocation(router.query.assetId as string)
  }, [router.query.assetId])
  useEffect(() => {
    fetchCenter(router.query.assetId as string)
  }, [router.query.assetId])
  return (
    <>
      <LoadScript
        googleMapsApiKey={process.env.NEXT_PUBLIC_GOOGLE_MAPS_API_KEY}
      >
        {center.map((center) => (
          tmpcenlat = +center.latitude,
          tmpcenlong = +center.longitude,
          <GoogleMap {...defaultMapConfig} center={{ lat: tmpcenlat, lng: tmpcenlong }} zoom={8}>
            {location.map((location) => (
              tmplat = +location.latitude,
              tmplong = +location.longitude,
                <Marker position={{ lat: tmplat, lng: tmplong }} draggable={false} />
            ))}
          {/* <Marker position={{ lat: 7.494726, lng: 102.612209 }} draggable={false} />
          <Marker position={{ lat: 8.263407, lng: 102.519507 }} draggable={false} />
          <Marker position={{ lat: 8.147704, lng: 102.255313 }} draggable={false} /> */}

        </GoogleMap>
        ))}

        {/* {center.map((center) => (
          <>
          <div> Center lat : {center.latitude} </div>
          <div> Center long : {center.longitude} </div>
          <div> Lat : {location.latitude} </div>
          <div> Long : {location.longitude} </div>
          </>
        ))}
        test */}
      </LoadScript>
    </>
  );
}
