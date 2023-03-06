// https://github.com/leixiaohua1020/simplest_encoder/blob/master/simplest_x264_encoder/simplest_x264_encoder.cpp
package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/x264-go/libx264"
	"github.com/moonfdd/x264-go/libx264common"
)

func main0() ffcommon.FInt {

	var ret ffcommon.FInt
	var y_size ffcommon.FInt
	var i, j ffcommon.FInt

	//FILE* fp_src  = fopen("../cuc_ieschool_640x360_yuv444p.yuv", "rb");
	fp_src, _ := os.Open("./resources/640x360_yuv420p.yuv")
	fp_dst_file := "./out/640x360_yuv420p.h264"
	fp_dst, _ := os.Create(fp_dst_file)

	//Encode 50 frame
	//if set 0, encode all frame
	var frame_num ffcommon.FInt = 0
	var csp ffcommon.FInt = libx264.X264_CSP_I420
	var width, height ffcommon.FInt = 640, 360

	var iNal ffcommon.FInt = 0
	var pNals *libx264.X264NalT
	var pHandle *libx264.X264T
	pPic_in := new(libx264.X264PictureT)
	pPic_out := new(libx264.X264PictureT)
	pParam := new(libx264.X264ParamT)

	//Check
	if fp_src == nil || fp_dst == nil {
		fmt.Printf("Error open files.\n")
		return -1
	}

	pParam.X264ParamDefault()
	pParam.IWidth = width
	pParam.IHeight = height
	/*
		//Param
		pParam->i_log_level  = X264_LOG_DEBUG;
		pParam->i_threads  = X264_SYNC_LOOKAHEAD_AUTO;
		pParam->i_frame_total = 0;
		pParam->i_keyint_max = 10;
		pParam->i_bframe  = 5;
		pParam->b_open_gop  = 0;
		pParam->i_bframe_pyramid = 0;
		pParam->rc.i_qp_constant=0;
		pParam->rc.i_qp_max=0;
		pParam->rc.i_qp_min=0;
		pParam->i_bframe_adaptive = X264_B_ADAPT_TRELLIS;
		pParam->i_fps_den  = 1;
		pParam->i_fps_num  = 25;
		pParam->i_timebase_den = pParam->i_fps_num;
		pParam->i_timebase_num = pParam->i_fps_den;
	*/
	pParam.ICsp = csp
	pParam.X264ParamApplyProfile(libx264.X264ProfileNames[5])

	pHandle = pParam.X264EncoderOpen164()

	pPic_out.X264PictureInit()
	pPic_in.X264PictureAlloc(csp, pParam.IWidth, pParam.IHeight)

	//ret = x264_encoder_headers(pHandle, &pNals, &iNal);

	y_size = pParam.IWidth * pParam.IHeight
	//detect frame number
	if frame_num == 0 {
		fi, _ := fp_src.Stat()
		switch csp {
		case libx264.X264_CSP_I444:
			frame_num = int32(fi.Size()) / (y_size * 3)
		case libx264.X264_CSP_I420:
			frame_num = int32(fi.Size()) / (y_size * 3 / 2)
		default:
			fmt.Printf("Colorspace Not Support.\n")
			return -1
		}
	}

	//Loop to Encode
	for i = 0; i < frame_num; i++ {
		switch csp {
		case libx264.X264_CSP_I444:

			fp_src.Read(ffcommon.ByteSliceFromByteP(pPic_in.Img.Plane[0], int(y_size))) //Y
			fp_src.Read(ffcommon.ByteSliceFromByteP(pPic_in.Img.Plane[1], int(y_size))) //U
			fp_src.Read(ffcommon.ByteSliceFromByteP(pPic_in.Img.Plane[2], int(y_size))) //V

		case libx264.X264_CSP_I420:

			fp_src.Read(ffcommon.ByteSliceFromByteP(pPic_in.Img.Plane[0], int(y_size)))   //Y
			fp_src.Read(ffcommon.ByteSliceFromByteP(pPic_in.Img.Plane[1], int(y_size/4))) //U
			fp_src.Read(ffcommon.ByteSliceFromByteP(pPic_in.Img.Plane[2], int(y_size/4))) //V

		default:

			fmt.Printf("Colorspace Not Support.\n")
			return -1

		}
		pPic_in.IPts = int64(i)

		ret = pHandle.X264EncoderEncode(&pNals, &iNal, pPic_in, pPic_out)
		if ret < 0 {
			fmt.Printf("Error.\n")
			return -1
		}

		fmt.Printf("Succeed encode frame: %5d\n", i)

		for j = 0; j < iNal; j++ {
			a := unsafe.Sizeof(libx264.X264NalT{})
			pNal := (*libx264.X264NalT)(unsafe.Pointer(uintptr(unsafe.Pointer(pNals)) + uintptr(a*uintptr(j))))
			fp_dst.Write(ffcommon.ByteSliceFromByteP(pNal.PPayload, int(pNal.IPayload)))
		}
	}
	i = 0
	//flush encoder
	for {
		ret = pHandle.X264EncoderEncode(&pNals, &iNal, nil, pPic_out)
		if ret == 0 {
			break
		}
		fmt.Printf("Flush 1 frame.\n")
		for j = 0; j < iNal; j++ {
			a := unsafe.Sizeof(libx264.X264NalT{})
			pNal := (*libx264.X264NalT)(unsafe.Pointer(uintptr(unsafe.Pointer(pNals)) + uintptr(a*uintptr(j))))
			fp_dst.Write(ffcommon.ByteSliceFromByteP(pNal.PPayload, int(pNal.IPayload)))
		}
		i++
	}
	pPic_in.X264PictureClean()
	pHandle.X264EncoderClose()
	pHandle = nil

	fp_src.Close()
	fp_dst.Close()

	fmt.Printf("\nffplay %s\n", fp_dst_file)

	return 0
}

func main() {
	fmt.Println(libx264.X264_POINTVER)
	os.Setenv("Path", os.Getenv("Path")+";./lib")
	libx264common.SetLibx264Path("./lib/libx264-164.dll")
	main0()
}
