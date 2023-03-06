package libx264

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/x264-go/libx264common"
)

/*****************************************************************************
 * x264.h: x264 public header
 *****************************************************************************
 * Copyright (C) 2003-2023 x264 project
 *
 * Authors: Laurent Aimar <fenrir@via.ecp.fr>
 *          Loren Merritt <lorenm@u.washington.edu>
 *          Fiona Glaser <fiona@x264.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02111, USA.
 *
 * This program is also available under a commercial proprietary license.
 * For more information, contact us at licensing@x264.com.
 *****************************************************************************/

// #ifndef X264_X264_H
// const X264_X264_H

// #ifdef __cplusplus
// extern "C" {
// #endif

// #if !defined(_STDINT_H) && !defined(_STDINT_H_) && !defined(_STDINT_H_INCLUDED) && !defined(_STDINT) &&\
//     !defined(_SYS_STDINT_H_) && !defined(_INTTYPES_H) && !defined(_INTTYPES_H_) && !defined(_INTTYPES)
// # ifdef _MSC_VER
// #  pragma message("You must include stdint.h or inttypes.h before x264.h")
// # else
// #  warning You must include stdint.h or inttypes.h before x264.h
// # endif
// #endif

// #include <stdarg.h>

// #include "x264_config.h"

const X264_BUILD = 164

// #ifdef _WIN32
// #   define X264_DLL_IMPORT __declspec(dllimport)
// #   define X264_DLL_EXPORT __declspec(dllexport)
// #else
// #   if defined(__GNUC__) && (__GNUC__ >= 4)
// #       define X264_DLL_IMPORT
// #       define X264_DLL_EXPORT __attribute__((visibility("default")))
// #   else
// #       define X264_DLL_IMPORT
// #       define X264_DLL_EXPORT
// #   endif
// #endif

/* Application developers planning to link against a shared library version of
 * libx264 from a Microsoft Visual Studio or similar development environment
 * will need to define X264_API_IMPORTS before including this header.
 * This clause does not apply to MinGW, similar development environments, or non
 * Windows platforms. */
// #ifdef X264_API_IMPORTS
// #   define X264_API X264_DLL_IMPORT
// #else
// #   ifdef X264_API_EXPORTS
// #       define X264_API X264_DLL_EXPORT
// #   else
// #       define X264_API
// #   endif
// #endif

/* x264_t:
 *      opaque handler for encoder */
// typedef struct x264_t x264_t;
type X264T struct{}

/****************************************************************************
 * NAL structure and functions
 ****************************************************************************/
type NalUnitTypeE ffcommon.FEnum

const (
	NAL_UNKNOWN   = 0
	NAL_SLICE     = 1
	NAL_SLICE_DPA = 2
	NAL_SLICE_DPB = 3
	NAL_SLICE_DPC = 4
	NAL_SLICE_IDR = 5 /* ref_idc != 0 */
	NAL_SEI       = 6 /* ref_idc == 0 */
	NAL_SPS       = 7
	NAL_PPS       = 8
	NAL_AUD       = 9
	NAL_FILLER    = 12
	/* ref_idc == 0 for 6,9,10,11,12 */
)

type NalPriorityE ffcommon.FEnum

const (
	NAL_PRIORITY_DISPOSABLE = 0
	NAL_PRIORITY_LOW        = 1
	NAL_PRIORITY_HIGH       = 2
	NAL_PRIORITY_HIGHEST    = 3
)

// /* The data within the payload is already NAL-encapsulated; the ref_idc and type
//   - are merely in the struct for easy access by the calling application.
//   - All data returned in an x264_nal_t, including the data in p_payload, is no longer
//   - valid after the next call to x264_encoder_encode.  Thus it must be used or copied
//   - before calling x264_encoder_encode or x264_encoder_headers again. */
type X264NalT struct {
	IRefIdc        ffcommon.FInt /* nal_priority_e */
	IType          ffcommon.FInt /* nal_unit_type_e */
	BLongStartcode ffcommon.FInt
	IFirstMb       ffcommon.FInt /* If this NAL is a slice, the index of the first MB in the slice. */
	ILastMb        ffcommon.FInt /* If this NAL is a slice, the index of the last MB in the slice. */

	/* Size of payload (including any padding) in bytes. */
	IPayload ffcommon.FInt
	/* If param->b_annexb is set, Annex-B bytestream with startcode.
	 * Otherwise, startcode is replaced with a 4-byte size.
	 * This size is the size used in mp4/similar muxing; it is equal to i_payload-4 */
	PPayload *ffcommon.FUint8T

	/* Size of padding in bytes. */
	IPadding ffcommon.FInt
}

/****************************************************************************
 * Encoder parameters
 ****************************************************************************/
/* CPU flags */

/* x86 */
const X264_CPU_MMX = (1 << 0)
const X264_CPU_MMX2 = (1 << 1) /* MMX2 aka MMXEXT aka ISSE */
const X264_CPU_MMXEXT = X264_CPU_MMX2
const X264_CPU_SSE = (1 << 2)
const X264_CPU_SSE2 = (1 << 3)
const X264_CPU_LZCNT = (1 << 4)
const X264_CPU_SSE3 = (1 << 5)
const X264_CPU_SSSE3 = (1 << 6)
const X264_CPU_SSE4 = (1 << 7)  /* SSE4.1 */
const X264_CPU_SSE42 = (1 << 8) /* SSE4.2 */
const X264_CPU_AVX = (1 << 9)   /* Requires OS support even if YMM registers aren't used */
const X264_CPU_XOP = (1 << 10)  /* AMD XOP */
const X264_CPU_FMA4 = (1 << 11) /* AMD FMA4 */
const X264_CPU_FMA3 = (1 << 12)
const X264_CPU_BMI1 = (1 << 13)
const X264_CPU_BMI2 = (1 << 14)
const X264_CPU_AVX2 = (1 << 15)
const X264_CPU_AVX512 = (1 << 16) /* AVX-512 {F, CD, BW, DQ, VL}, requires OS support */
/* x86 modifiers */
const X264_CPU_CACHELINE_32 = (1 << 17) /* avoid memory loads that span the border between two cachelines */
const X264_CPU_CACHELINE_64 = (1 << 18) /* 32/64 is the size of a cacheline in bytes */
const X264_CPU_SSE2_IS_SLOW = (1 << 19) /* avoid most SSE2 functions on Athlon64 */
const X264_CPU_SSE2_IS_FAST = (1 << 20) /* a few functions are only faster on Core2 and Phenom */
const X264_CPU_SLOW_SHUFFLE = (1 << 21) /* The Conroe has a slow shuffle unit (relative to overall SSE performance) */
const X264_CPU_STACK_MOD4 = (1 << 22)   /* if stack is only mod4 and not mod16 */
const X264_CPU_SLOW_ATOM = (1 << 23)    /* The Atom is terrible: slow SSE unaligned loads, slow
 * SIMD multiplies, slow SIMD variable shifts, slow pshufb,
 * cacheline split penalties -- gather everything here that
 * isn't shared by other CPUs to avoid making half a dozen
 * new SLOW flags. */
const X264_CPU_SLOW_PSHUFB = (1 << 24)  /* such as on the Intel Atom */
const X264_CPU_SLOW_PALIGNR = (1 << 25) /* such as on the AMD Bobcat */

/* PowerPC */
const X264_CPU_ALTIVEC = 0x0000001

/* ARM and AArch64 */
const X264_CPU_ARMV6 = 0x0000001
const X264_CPU_NEON = 0x0000002          /* ARM NEON */
const X264_CPU_FAST_NEON_MRC = 0x0000004 /* Transfer from NEON to ARM register is fast (Cortex-A9) */
const X264_CPU_ARMV8 = 0x0000008

/* MIPS */
const X264_CPU_MSA = 0x0000001 /* MIPS MSA */

/* Analyse flags */
const X264_ANALYSE_I4x4 = 0x0001      /* Analyse i4x4 */
const X264_ANALYSE_I8x8 = 0x0002      /* Analyse i8x8 (requires 8x8 transform) */
const X264_ANALYSE_PSUB16x16 = 0x0010 /* Analyse p16x8, p8x16 and p8x8 */
const X264_ANALYSE_PSUB8x8 = 0x0020   /* Analyse p8x4, p4x8, p4x4 */
const X264_ANALYSE_BSUB16x16 = 0x0100 /* Analyse b16x8, b8x16 and b8x8 */

const X264_DIRECT_PRED_NONE = 0
const X264_DIRECT_PRED_SPATIAL = 1
const X264_DIRECT_PRED_TEMPORAL = 2
const X264_DIRECT_PRED_AUTO = 3
const X264_ME_DIA = 0
const X264_ME_HEX = 1
const X264_ME_UMH = 2
const X264_ME_ESA = 3
const X264_ME_TESA = 4
const X264_CQM_FLAT = 0
const X264_CQM_JVT = 1
const X264_CQM_CUSTOM = 2
const X264_RC_CQP = 0
const X264_RC_CRF = 1
const X264_RC_ABR = 2
const X264_QP_AUTO = 0
const X264_AQ_NONE = 0
const X264_AQ_VARIANCE = 1
const X264_AQ_AUTOVARIANCE = 2
const X264_AQ_AUTOVARIANCE_BIASED = 3
const X264_B_ADAPT_NONE = 0
const X264_B_ADAPT_FAST = 1
const X264_B_ADAPT_TRELLIS = 2
const X264_WEIGHTP_NONE = 0
const X264_WEIGHTP_SIMPLE = 1
const X264_WEIGHTP_SMART = 2
const X264_B_PYRAMID_NONE = 0
const X264_B_PYRAMID_STRICT = 1
const X264_B_PYRAMID_NORMAL = 2
const X264_KEYINT_MIN_AUTO = 0
const X264_KEYINT_MAX_INFINITE = (1 << 30)

/* AVC-Intra flavors */
const X264_AVCINTRA_FLAVOR_PANASONIC = 0
const X264_AVCINTRA_FLAVOR_SONY = 1

// static const char * const x264_direct_pred_names[] = { "none", "spatial", "temporal", "auto", 0 };
var X264DirectPredNames = []string{"none", "spatial", "temporal", "auto"}

// static const char * const x264_motion_est_names[] = { "dia", "hex", "umh", "esa", "tesa", 0 };
var X264MotionEstNames = []string{"dia", "hex", "umh", "esa", "tesa"}

// static const char * const x264_b_pyramid_names[] = { "none", "strict", "normal", 0 };
var X264BPyramidNames = []string{"none", "strict", "normal"}

// static const char * const x264_overscan_names[] = { "undef", "show", "crop", 0 };
var X264OverscanNames = []string{"undef", "show", "crop"}

// static const char * const x264_vidformat_names[] = { "component", "pal", "ntsc", "secam", "mac", "undef", 0 };
var X264VidformatNames = []string{"component", "pal", "ntsc", "secam", "mac", "undef"}

// static const char * const x264_fullrange_names[] = { "off", "on", 0 };
var X264FullrangeNames = []string{"off", "on"}

// static const char * const x264_colorprim_names[] = { "", "bt709", "undef", "", "bt470m", "bt470bg", "smpte170m", "smpte240m", "film", "bt2020", "smpte428",
//
//	"smpte431", "smpte432", 0 };
var X264ColorprimNames = []string{"", "bt709", "undef", "", "bt470m", "bt470bg", "smpte170m", "smpte240m", "film", "bt2020", "smpte428", "smpte431", "smpte432"}

// static const char * const x264_transfer_names[] = { "", "bt709", "undef", "", "bt470m", "bt470bg", "smpte170m", "smpte240m", "linear", "log100", "log316",
//
//	"iec61966-2-4", "bt1361e", "iec61966-2-1", "bt2020-10", "bt2020-12", "smpte2084", "smpte428", "arib-std-b67", 0 };
var X264TransferNames = []string{"", "bt709", "undef", "", "bt470m", "bt470bg", "smpte170m", "smpte240m", "linear", "log100", "log316", "iec61966-2-4", "bt1361e", "iec61966-2-1", "bt2020-10", "bt2020-12", "smpte2084", "smpte428", "arib-std-b67"}

// static const char * const x264_colmatrix_names[] = { "GBR", "bt709", "undef", "", "fcc", "bt470bg", "smpte170m", "smpte240m", "YCgCo", "bt2020nc", "bt2020c",
//
//	"smpte2085", "chroma-derived-nc", "chroma-derived-c", "ICtCp", 0 };
var X264ColmatrixNames = []string{"GBR", "bt709", "undef", "", "fcc", "bt470bg", "smpte170m", "smpte240m", "YCgCo", "bt2020nc", "bt2020c", "smpte2085", "chroma-derived-nc", "chroma-derived-c", "ICtCp"}

// static const char * const x264_nal_hrd_names[] = { "none", "vbr", "cbr", 0 };
var X264NalHrdNames = []string{"none", "vbr", "cbr"}

// static const char * const x264_avcintra_flavor_names[] = { "panasonic", "sony", 0 };
var X264AvcintraFlavorNames = []string{"panasonic", "sony"}

/* Colorspace type */
const X264_CSP_MASK = 0x00ff       /* */
const X264_CSP_NONE = 0x0000       /* Invalid mode     */
const X264_CSP_I400 = 0x0001       /* monochrome 4:0:0 */
const X264_CSP_I420 = 0x0002       /* yuv 4:2:0 planar */
const X264_CSP_YV12 = 0x0003       /* yvu 4:2:0 planar */
const X264_CSP_NV12 = 0x0004       /* yuv 4:2:0, with one y plane and one packed u+v */
const X264_CSP_NV21 = 0x0005       /* yuv 4:2:0, with one y plane and one packed v+u */
const X264_CSP_I422 = 0x0006       /* yuv 4:2:2 planar */
const X264_CSP_YV16 = 0x0007       /* yvu 4:2:2 planar */
const X264_CSP_NV16 = 0x0008       /* yuv 4:2:2, with one y plane and one packed u+v */
const X264_CSP_YUYV = 0x0009       /* yuyv 4:2:2 packed */
const X264_CSP_UYVY = 0x000a       /* uyvy 4:2:2 packed */
const X264_CSP_V210 = 0x000b       /* 10-bit yuv 4:2:2 packed in 32 */
const X264_CSP_I444 = 0x000c       /* yuv 4:4:4 planar */
const X264_CSP_YV24 = 0x000d       /* yvu 4:4:4 planar */
const X264_CSP_BGR = 0x000e        /* packed bgr 24bits */
const X264_CSP_BGRA = 0x000f       /* packed bgr 32bits */
const X264_CSP_RGB = 0x0010        /* packed rgb 24bits */
const X264_CSP_MAX = 0x0011        /* end of list */
const X264_CSP_VFLIP = 0x1000      /* the csp is vertically flipped */
const X264_CSP_HIGH_DEPTH = 0x2000 /* the csp has a depth of 16 bits per pixel component */

/* Slice type */
const X264_TYPE_AUTO = 0x0000 /* Let x264 choose the right type */
const X264_TYPE_IDR = 0x0001
const X264_TYPE_I = 0x0002
const X264_TYPE_P = 0x0003
const X264_TYPE_BREF = 0x0004 /* Non-disposable B-frame */
const X264_TYPE_B = 0x0005
const X264_TYPE_KEYFRAME = 0x0006 /* IDR or I depending on b_open_gop option */
// const IS_X264_TYPE_I(x) ((x)==X264_TYPE_I || (x)==X264_TYPE_IDR || (x)==X264_TYPE_KEYFRAME)
// const IS_X264_TYPE_B(x) ((x)==X264_TYPE_B || (x)==X264_TYPE_BREF)

/* Log level */
const X264_LOG_NONE = (-1)
const X264_LOG_ERROR = 0
const X264_LOG_WARNING = 1
const X264_LOG_INFO = 2
const X264_LOG_DEBUG = 3

/* Threading */
const X264_THREADS_AUTO = 0           /* Automatically select optimal number of threads */
const X264_SYNC_LOOKAHEAD_AUTO = (-1) /* Automatically select optimal lookahead thread buffer size */

/* HRD */
const X264_NAL_HRD_NONE = 0
const X264_NAL_HRD_VBR = 1
const X264_NAL_HRD_CBR = 2

/* Zones: override ratecontrol or other options for specific sections of the video.
 * See x264_encoder_reconfig() for which options can be changed.
 * If zones overlap, whichever comes later in the list takes precedence. */
type X264ZoneT struct {
	IStart, IEnd   ffcommon.FInt /* range of frame numbers */
	BForceQp       ffcommon.FInt /* whether to use qp vs bitrate factor */
	IQp            ffcommon.FInt
	FBitrateFactor ffcommon.FFloat
	Param          *X264ParamT
}

type X264ParamT struct {

	/* CPU flags */
	cpu                 ffcommon.FUint32T
	i_threads           ffcommon.FInt /* encode multiple frames in parallel */
	i_lookahead_threads ffcommon.FInt /* multiple threads for lookahead analysis */
	b_sliced_threads    ffcommon.FInt /* Whether to use slice-based threading. */
	b_deterministic     ffcommon.FInt /* whether to allow non-deterministic optimizations when threaded */
	b_cpu_independent   ffcommon.FInt /* force canonical behavior rather than cpu-dependent optimal algorithms */
	i_sync_lookahead    ffcommon.FInt /* threaded lookahead buffer */

	/* Video Properties */
	IWidth        ffcommon.FInt
	IHeight       ffcommon.FInt
	ICsp          ffcommon.FInt /* CSP of encoded bitstream */
	i_bitdepth    ffcommon.FInt
	i_level_idc   ffcommon.FInt
	i_frame_total ffcommon.FInt /* number of frames to encode if known, else 0 */

	/* NAL HRD
	 * Uses Buffering and Picture Timing SEIs to signal HRD
	 * The HRD in H.264 was not designed with VFR in mind.
	 * It is therefore not recommendeded to use NAL HRD with VFR.
	 * Furthermore, reconfiguring the VBV (via x264_encoder_reconfig)
	 * will currently generate invalid HRD. */
	i_nal_hrd ffcommon.FInt

	vui struct {
		/* they will be reduced to be 0 < x <= 65535 and prime */
		i_sar_height ffcommon.FInt
		i_sar_width  ffcommon.FInt

		i_overscan ffcommon.FInt /* 0=undef, 1=no overscan, 2=overscan */

		/* see h264 annex E for the values of the following */
		i_vidformat  ffcommon.FInt
		b_fullrange  ffcommon.FInt
		i_colorprim  ffcommon.FInt
		i_transfer   ffcommon.FInt
		i_colmatrix  ffcommon.FInt
		i_chroma_loc ffcommon.FInt /* both top & bottom */
	}

	/* Bitstream parameters */
	i_frame_reference ffcommon.FInt /* Maximum number of reference frames */
	i_dpb_size        ffcommon.FInt /* Force a DPB size larger than that implied by B-frames and reference frames.
	 * Useful in combination with interactive error resilience. */
	i_keyint_max         ffcommon.FInt /* Force an IDR keyframe at this interval */
	i_keyint_min         ffcommon.FInt /* Scenecuts closer together than this are coded as I, not IDR. */
	i_scenecut_threshold ffcommon.FInt /* how aggressively to insert extra I frames */
	b_intra_refresh      ffcommon.FInt /* Whether or not to use periodic intra refresh instead of IDR frames. */

	i_bframe          ffcommon.FInt /* how many b-frame between 2 references pictures */
	i_bframe_adaptive ffcommon.FInt
	i_bframe_bias     ffcommon.FInt
	i_bframe_pyramid  ffcommon.FInt /* Keep some B-frames as references: 0=off, 1=strict hierarchical, 2=normal */
	b_open_gop        ffcommon.FInt
	b_bluray_compat   ffcommon.FInt
	i_avcintra_class  ffcommon.FInt
	i_avcintra_flavor ffcommon.FInt

	b_deblocking_filter         ffcommon.FInt
	i_deblocking_filter_alphac0 ffcommon.FInt /* [-6, 6] -6 light filter, 6 strong */
	i_deblocking_filter_beta    ffcommon.FInt /* [-6, 6]  idem */

	b_cabac          ffcommon.FInt
	i_cabac_init_idc ffcommon.FInt

	b_interlaced        ffcommon.FInt
	b_constrained_intra ffcommon.FInt

	i_cqm_preset ffcommon.FInt
	psz_cqm_file ffcommon.FCharPStruct /* filename (in UTF-8) of CQM file, JM format */
	cqm_4iy      [16]ffcommon.FUint8T  /* used only if i_cqm_preset == X264_CQM_CUSTOM */
	cqm_4py      [16]ffcommon.FUint8T
	cqm_4ic      [16]ffcommon.FUint8T
	cqm_4pc      [16]ffcommon.FUint8T
	cqm_8iy      [64]ffcommon.FUint8T
	cqm_8py      [64]ffcommon.FUint8T
	cqm_8ic      [64]ffcommon.FUint8T
	cqm_8pc      [64]ffcommon.FUint8T

	/* Log */
	//void        (*pf_log)( void *, int i_level, const char *psz, va_list );
	pf_log        uintptr
	p_log_private ffcommon.FVoidP
	i_log_level   ffcommon.FInt
	b_full_recon  ffcommon.FInt         /* fully reconstruct frames, even when not necessary for encoding.  Implied by psz_dump_yuv */
	psz_dump_yuv  ffcommon.FCharPStruct /* filename (in UTF-8) for reconstructed frames */

	/* Encoder analyser parameters */
	analyse struct {
		intra ffcommon.FUnsignedInt /* intra partitions */
		inter ffcommon.FUnsignedInt /* inter partitions */

		b_transform_8x8    ffcommon.FInt
		i_weighted_pred    ffcommon.FInt /* weighting for P-frames */
		b_weighted_bipred  ffcommon.FInt /* implicit weighting for B-frames */
		i_direct_mv_pred   ffcommon.FInt /* spatial vs temporal mv prediction */
		i_chroma_qp_offset ffcommon.FInt

		i_me_method        ffcommon.FInt   /* motion estimation algorithm to use (X264_ME_*) */
		i_me_range         ffcommon.FInt   /* integer pixel motion estimation search range (from predicted mv) */
		i_mv_range         ffcommon.FInt   /* maximum length of a mv (in pixels). -1 = auto, based on level */
		i_mv_range_thread  ffcommon.FInt   /* minimum space between threads. -1 = auto, based on number of threads. */
		i_subpel_refine    ffcommon.FInt   /* subpixel motion estimation quality */
		b_chroma_me        ffcommon.FInt   /* chroma ME for subpel and mode decision in P-frames */
		b_mixed_references ffcommon.FInt   /* allow each mb partition to have its own reference number */
		i_trellis          ffcommon.FInt   /* trellis RD quantization */
		b_fast_pskip       ffcommon.FInt   /* early SKIP detection on P-frames */
		b_dct_decimate     ffcommon.FInt   /* transform coefficient thresholding on P-frames */
		i_noise_reduction  ffcommon.FInt   /* adaptive pseudo-deadzone */
		f_psy_rd           ffcommon.FFloat /* Psy RD strength */
		f_psy_trellis      ffcommon.FFloat /* Psy trellis strength */
		b_psy              ffcommon.FInt   /* Toggle all psy optimizations */

		b_mb_info        ffcommon.FInt /* Use input mb_info data in x264_picture_t */
		b_mb_info_update ffcommon.FInt /* Update the values in mb_info according to the results of encoding. */

		/* the deadzone size that will be used in luma quantization */
		i_luma_deadzone [2]ffcommon.FInt /* {inter, intra} */

		b_psnr ffcommon.FInt /* compute and print PSNR stats */
		b_ssim ffcommon.FInt /* compute and print SSIM stats */
	}

	/* Rate control parameters */
	rc struct {
		i_rc_method ffcommon.FInt /* X264_RC_* */

		i_qp_constant ffcommon.FInt /* 0=lossless */
		i_qp_min      ffcommon.FInt /* min allowed QP value */
		i_qp_max      ffcommon.FInt /* max allowed QP value */
		i_qp_step     ffcommon.FInt /* max QP step between frames */

		i_bitrate         ffcommon.FInt
		f_rf_constant     ffcommon.FFloat /* 1pass VBR, nominal QP */
		f_rf_constant_max ffcommon.FFloat /* In CRF mode, maximum CRF as caused by VBV */
		f_rate_tolerance  ffcommon.FFloat
		i_vbv_max_bitrate ffcommon.FInt
		i_vbv_buffer_size ffcommon.FInt
		f_vbv_buffer_init ffcommon.FFloat /* <=1: fraction of buffer_size. >1: kbit */
		f_ip_factor       ffcommon.FFloat
		f_pb_factor       ffcommon.FFloat

		/* VBV filler: force CBR VBV and use filler bytes to ensure hard-CBR.
		 * Implied by NAL-HRD CBR. */
		b_filler ffcommon.FInt

		i_aq_mode     ffcommon.FInt /* psy adaptive QP. (X264_AQ_*) */
		f_aq_strength ffcommon.FFloat
		b_mb_tree     ffcommon.FInt /* Macroblock-tree ratecontrol. */
		i_lookahead   ffcommon.FInt

		/* 2pass */
		b_stat_write ffcommon.FInt         /* Enable stat writing in psz_stat_out */
		psz_stat_out ffcommon.FCharPStruct /* output filename (in UTF-8) of the 2pass stats file */
		b_stat_read  ffcommon.FInt         /* Read stat from psz_stat_in and use it */
		psz_stat_in  ffcommon.FCharPStruct /* input filename (in UTF-8) of the 2pass stats file */

		/* 2pass params (same as ffmpeg ones) */
		f_qcompress       ffcommon.FFloat       /* 0.0 => cbr, 1.0 => constant qp */
		f_qblur           ffcommon.FFloat       /* temporally blur quants */
		f_complexity_blur ffcommon.FFloat       /* temporally blur complexity */
		zones             *X264ZoneT            /* ratecontrol overrides */
		i_zones           ffcommon.FInt         /* number of zone_t's */
		psz_zones         ffcommon.FCharPStruct /* alternate method of specifying zones */
	}

	/* Cropping Rectangle parameters: added to those implicitly defined by
	   non-mod16 video resolutions. */
	crop_rect struct {
		i_left   ffcommon.FInt
		i_top    ffcommon.FInt
		i_right  ffcommon.FInt
		i_bottom ffcommon.FInt
	}

	/* frame packing arrangement flag */
	i_frame_packing ffcommon.FInt

	/* mastering display SEI: Primary and white point chromaticity coordinates
	   in 0.00002 increments. Brightness units are 0.0001 cd/m^2. */
	mastering_display struct {
		b_mastering_display ffcommon.FInt /* enable writing this SEI */
		i_green_x           ffcommon.FInt
		i_green_y           ffcommon.FInt
		i_blue_x            ffcommon.FInt
		i_blue_y            ffcommon.FInt
		i_red_x             ffcommon.FInt
		i_red_y             ffcommon.FInt
		i_white_x           ffcommon.FInt
		i_white_y           ffcommon.FInt
		i_display_max       ffcommon.FInt64T
		i_display_min       ffcommon.FInt64T
	}

	/* content light level SEI */
	content_light_level struct {
		b_cll      ffcommon.FInt /* enable writing this SEI */
		i_max_cll  ffcommon.FInt
		i_max_fall ffcommon.FInt
	}

	/* alternative transfer SEI */
	i_alternative_transfer ffcommon.FInt

	/* Muxing parameters */
	b_aud            ffcommon.FInt /* generate access unit delimiters */
	b_repeat_headers ffcommon.FInt /* put SPS/PPS before each keyframe */
	b_annexb         ffcommon.FInt /* if set, place start codes (4 bytes) before NAL units,
	 * otherwise place size (4 bytes) before NAL units. */
	i_sps_id    ffcommon.FInt /* SPS and PPS id number */
	b_vfr_input ffcommon.FInt /* VFR input.  If 1, use timebase and timestamps for ratecontrol purposes.
	 * If 0, use fps only. */
	b_pulldown     ffcommon.FInt /* use explicitly set timebase for CFR */
	i_fps_num      ffcommon.FUint32T
	i_fps_den      ffcommon.FUint32T
	i_timebase_num ffcommon.FUint32T /* Timebase numerator */
	i_timebase_den ffcommon.FUint32T /* Timebase denominator */

	b_tff ffcommon.FInt

	/* Pulldown:
	 * The correct pic_struct must be passed with each input frame.
	 * The input timebase should be the timebase corresponding to the output framerate. This should be constant.
	 * e.g. for 3:2 pulldown timebase should be 1001/30000
	 * The PTS passed with each frame must be the PTS of the frame after pulldown is applied.
	 * Frame doubling and tripling require b_vfr_input set to zero (see H.264 Table D-1)
	 *
	 * Pulldown changes are not clearly defined in H.264. Therefore, it is the calling app's responsibility to manage this.
	 */

	b_pic_struct ffcommon.FInt

	/* Fake Interlaced.
	 *
	 * Used only when b_interlaced=0. Setting this flag makes it possible to flag the stream as PAFF interlaced yet
	 * encode all frames progessively. It is useful for encoding 25p and 30p Blu-Ray streams.
	 */

	b_fake_interlaced ffcommon.FInt

	/* Don't optimize header parameters based on video content, e.g. ensure that splitting an input video, compressing
	 * each part, and stitching them back together will result in identical SPS/PPS. This is necessary for stitching
	 * with container formats that don't allow multiple SPS/PPS. */
	b_stitchable ffcommon.FInt

	b_opencl         ffcommon.FInt         /* use OpenCL when available */
	i_opencl_device  ffcommon.FInt         /* specify count of GPU devices to skip, for CLI users */
	opencl_device_id ffcommon.FVoidP       /* pass explicit cl_device_id as void*, for API users */
	psz_clbin_file   ffcommon.FCharPStruct /* filename (in UTF-8) of the compiled OpenCL kernel cache file */

	/* Slicing parameters */
	i_slice_max_size  ffcommon.FInt /* Max size per slice in bytes; includes estimated NAL overhead. */
	i_slice_max_mbs   ffcommon.FInt /* Max number of MBs per slice; overrides i_slice_count. */
	i_slice_min_mbs   ffcommon.FInt /* Min number of MBs per slice */
	i_slice_count     ffcommon.FInt /* Number of slices per frame: forces rectangular slices. */
	i_slice_count_max ffcommon.FInt /* Absolute cap on slices per frame; stops applying slice-max-size
	 * and slice-max-mbs if this is reached. */

	/* Optional callback for freeing this x264_param_t when it is done being used.
	 * Only used when the x264_param_t sits in memory for an indefinite period of time,
	 * i.e. when an x264_param_t is passed to x264_t in an x264_picture_t or in zones.
	 * Not used when x264_encoder_reconfig is called directly. */
	//void (*param_free)( void* );
	param_free uintptr

	/* Optional low-level callback for low-latency encoding.  Called for each output NAL unit
	 * immediately after the NAL unit is finished encoding.  This allows the calling application
	 * to begin processing video data (e.g. by sending packets over a network) before the frame
	 * is done encoding.
	 *
	 * This callback MUST do the following in order to work correctly:
	 * 1) Have available an output buffer of at least size nal->i_payload*3/2 + 5 + 64.
	 * 2) Call x264_nal_encode( h, dst, nal ), where dst is the output buffer.
	 * After these steps, the content of nal is valid and can be used in the same way as if
	 * the NAL unit were output by x264_encoder_encode.
	 *
	 * This does not need to be synchronous with the encoding process: the data pointed to
	 * by nal (both before and after x264_nal_encode) will remain valid until the next
	 * x264_encoder_encode call.  The callback must be re-entrant.
	 *
	 * This callback does not work with frame-based threads; threads must be disabled
	 * or sliced-threads enabled.  This callback also does not work as one would expect
	 * with HRD -- since the buffering period SEI cannot be calculated until the frame
	 * is finished encoding, it will not be sent via this callback.
	 *
	 * Note also that the NALs are not necessarily returned in order when sliced threads is
	 * enabled.  Accordingly, the variable i_first_mb and i_last_mb are available in
	 * x264_nal_t to help the calling application reorder the slices if necessary.
	 *
	 * When this callback is enabled, x264_encoder_encode does not return valid NALs;
	 * the calling application is expected to acquire all output NALs through the callback.
	 *
	 * It is generally sensible to combine this callback with a use of slice-max-mbs or
	 * slice-max-size.
	 *
	 * The opaque pointer is the opaque pointer from the input frame associated with this
	 * NAL unit. This helps distinguish between nalu_process calls from different sources,
	 * e.g. if doing multiple encodes in one process.
	 */
	//void (*nalu_process)( x264_t *h, x264_nal_t *nal, void *opaque );
	nalu_process uintptr

	/* For internal use only */
	opaque ffcommon.FVoidP
}

// X264_API void x264_nal_encode( x264_t *h, uint8_t *dst, x264_nal_t *nal );
func (h *X264T) X264NalEncode(dst *ffcommon.FUint8T, nal *X264NalT) {
	libx264common.GetLibx264Dll().NewProc("x264_nal_encode").Call(
		uintptr(unsafe.Pointer(h)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(nal)),
	)
}

// /****************************************************************************
//  * H.264 level restriction information
//  ****************************************************************************/

type X264LevelT struct {
	level_idc   ffcommon.FUint8T
	mbps        ffcommon.FInt32T  /* max macroblock processing rate (macroblocks/sec) */
	frame_size  ffcommon.FInt32T  /* max frame size (macroblocks) */
	dpb         ffcommon.FInt32T  /* max decoded picture buffer (mbs) */
	bitrate     ffcommon.FInt32T  /* max bitrate (kbit/sec) */
	cpb         ffcommon.FInt32T  /* max vbv buffer (kbit) */
	mv_range    ffcommon.FUint16T /* max vertical mv component range (pixels) */
	mvs_per_2mb ffcommon.FUint8T  /* max mvs per 2 consecutive mbs. */
	slice_rate  ffcommon.FUint8T  /* ?? */
	mincr       ffcommon.FUint8T  /* min compression ratio */
	bipred8x8   ffcommon.FUint8T  /* limit bipred to >=8x8 */
	direct8x8   ffcommon.FUint8T  /* limit b_direct to >=8x8 */
	frame_only  ffcommon.FUint8T  /* forbid interlacing */
}

/* all of the levels defined in the standard, terminated by .level_idc=0 */
// X264_API extern const x264_level_t x264_levels[];

/****************************************************************************
 * Basic parameter handling functions
 ****************************************************************************/

/* x264_param_default:
 *      fill x264_param_t with default values and do CPU detection */
//X264_API void x264_param_default( x264_param_t * );
func (this *X264ParamT) X264ParamDefault() {
	libx264common.GetLibx264Dll().NewProc("x264_param_default").Call(
		uintptr(unsafe.Pointer(this)),
	)
}

/* x264_param_parse:
 *  set one parameter by name.
 *  returns 0 on success, or returns one of the following errors.
 *  note: BAD_VALUE occurs only if it can't even parse the value,
 *  numerical range is not checked until x264_encoder_open() or
 *  x264_encoder_reconfig().
 *  value=NULL means "true" for boolean options, but is a BAD_VALUE for non-booleans.
 *  can allocate memory which should be freed by call of x264_param_cleanup. */
const X264_PARAM_BAD_NAME = (-1)
const X264_PARAM_BAD_VALUE = (-2)
const X264_PARAM_ALLOC_FAILED = (-3)

// X264_API int x264_param_parse( x264_param_t *, const char *name, const char *value );
func (this *X264ParamT) X264ParamParse(name, value ffcommon.FConstCharP) (res ffcommon.FInt) {
	t, _, _ := libx264common.GetLibx264Dll().NewProc("x264_param_parse").Call(
		uintptr(unsafe.Pointer(this)),
		ffcommon.UintPtrFromString(name),
		ffcommon.UintPtrFromString(value),
	)
	res = ffcommon.FInt(t)
	return
}

// /* x264_param_cleanup:
//   - Cleans up and frees allocated members of x264_param_t.
//   - This *does not* free the x264_param_t itself, as it may exist on the
//   - stack. It only frees any members of the struct that were allocated by
//   - x264 itself, in e.g. x264_param_parse(). */
//
// X264_API void x264_param_cleanup( x264_param_t *param );
func (param *X264ParamT) X264ParamCleanup() {
	libx264common.GetLibx264Dll().NewProc("x264_param_cleanup").Call(
		uintptr(unsafe.Pointer(param)),
	)
}

/****************************************************************************
 * Advanced parameter handling functions
 ****************************************************************************/

/* These functions expose the full power of x264's preset-tune-profile system for
 * easy adjustment of large numbers of internal parameters.
 *
 * In order to replicate x264CLI's option handling, these functions MUST be called
 * in the following order:
 * 1) x264_param_default_preset
 * 2) Custom user options (via param_parse or directly assigned variables)
 * 3) x264_param_apply_fastfirstpass
 * 4) x264_param_apply_profile
 *
 * Additionally, x264CLI does not apply step 3 if the preset chosen is "placebo"
 * or --slow-firstpass is set. */

/* x264_param_default_preset:
 *      The same as x264_param_default, but also use the passed preset and tune
 *      to modify the default settings.
 *      (either can be NULL, which implies no preset or no tune, respectively)
 *
 *      Currently available presets are, ordered from fastest to slowest: */
// static const char * const x264_preset_names[] = { "ultrafast", "superfast", "veryfast", "faster", "fast", "medium", "slow", "slower", "veryslow", "placebo", 0 };
var X264PresetNames = []string{"ultrafast", "superfast", "veryfast", "faster", "fast", "medium", "slow", "slower", "veryslow", "placebo"}

/*      The presets can also be indexed numerically, as in:
 *      x264_param_default_preset( &param, "3", ... )
 *      with ultrafast mapping to "0" and placebo mapping to "9".  This mapping mayX264_API
 *      of course change if new presets are added in between, but will always be
 *      ordered from fastest to slowest.
 *
 *      Warning: the speed of these presets scales dramatically.  Ultrafast is a full
 *      100 times faster than placebo!
 *
 *      Currently available tunings are: */
// static const char * const x264_tune_names[] = { "film", "animation", "grain", "stillimage", "psnr", "ssim", "fastdecode", "zerolatency", 0 };
var X264TuneNames = []string{"film", "animation", "grain", "stillimage", "psnr", "ssim", "fastdecode", "zerolatency"}

/*      Multiple tunings can be used if separated by a delimiter in ",./-+",
 *      however multiple psy tunings cannot be used.
 *      film, animation, grain, stillimage, psnr, and ssim are psy tunings.
 *
 *      returns 0 on success, negative on failure (e.g. invalid preset/tune name). */
// X264_API int x264_param_default_preset( x264_param_t *, const char *preset, const char *tune );
func (this *X264ParamT) X264ParamDefaultPreset(preset, tune ffcommon.FConstCharP) (res ffcommon.FInt) {
	t, _, _ := libx264common.GetLibx264Dll().NewProc("x264_param_default_preset").Call(
		uintptr(unsafe.Pointer(this)),
		ffcommon.UintPtrFromString(preset),
		ffcommon.UintPtrFromString(tune),
	)
	res = ffcommon.FInt(t)
	return
}

// /* x264_param_apply_fastfirstpass:
//   - If first-pass mode is set (rc.b_stat_read == 0, rc.b_stat_write == 1),
//   - modify the encoder settings to disable options generally not useful on
//   - the first pass. */
//
// X264_API void x264_param_apply_fastfirstpass( x264_param_t * );
func (this *X264ParamT) X264ParamApplyFastfirstpass() {
	libx264common.GetLibx264Dll().NewProc("x264_param_apply_fastfirstpass").Call(
		uintptr(unsafe.Pointer(this)),
	)
}

// /* x264_param_apply_profile:
//   - Applies the restrictions of the given profile.
//   - Currently available profiles are, from most to least restrictive: */
//
// static const char * const x264_profile_names[] = { "baseline", "main", "high", "high10", "high422", "high444", 0 };
var X264ProfileNames = []string{"baseline", "main", "high", "high10", "high422", "high444"}

// /*      (can be NULL, in which case the function will do nothing)
//
//	*
//	*      Does NOT guarantee that the given profile will be used: if the restrictions
//	*      of "High" are applied to settings that are already Baseline-compatible, the
//	*      stream will remain baseline.  In short, it does not increase settings, only
//	*      decrease them.
//	*
//	*      returns 0 on success, negative on failure (e.g. invalid profile name). */
//
// X264_API int x264_param_apply_profile( x264_param_t *, const char *profile );
func (this *X264ParamT) X264ParamApplyProfile(profile ffcommon.FConstCharP) (res ffcommon.FInt) {
	t, _, _ := libx264common.GetLibx264Dll().NewProc("x264_param_apply_profile").Call(
		uintptr(unsafe.Pointer(this)),
		ffcommon.UintPtrFromString(profile),
	)
	res = ffcommon.FInt(t)
	return
}

/****************************************************************************
 * Picture structures and functions
 ****************************************************************************/

/* x264_chroma_format:
 *      Specifies the chroma formats that x264 supports encoding. When this
 *      value is non-zero, then it represents a X264_CSP_* that is the only
 *      chroma format that x264 supports encoding. If the value is 0 then
 *      there are no restrictions. */
// X264_API extern const int x264_chroma_format;
type PicStructE ffcommon.FEnum

const (
	PIC_STRUCT_AUTO        = 0 // automatically decide (default)
	PIC_STRUCT_PROGRESSIVE = 1 // progressive frame
	// "TOP" and "BOTTOM" are not supported in x264 (PAFF only)
	PIC_STRUCT_TOP_BOTTOM        = 4 // top field followed by bottom
	PIC_STRUCT_BOTTOM_TOP        = 5 // bottom field followed by top
	PIC_STRUCT_TOP_BOTTOM_TOP    = 6 // top field, bottom field, top field repeated
	PIC_STRUCT_BOTTOM_TOP_BOTTOM = 7 // bottom field, top field, bottom field repeated
	PIC_STRUCT_DOUBLE            = 8 // double frame
	PIC_STRUCT_TRIPLE            = 9 // triple frame
)

type X264HrdT struct {
	cpb_initial_arrival_time ffcommon.FDouble
	cpb_final_arrival_time   ffcommon.FDouble
	cpb_removal_time         ffcommon.FDouble

	dpb_output_time ffcommon.FDouble
}

// /* Arbitrary user SEI:
//  * Payload size is in bytes and the payload pointer must be valid.
//  * Payload types and syntax can be found in Annex D of the H.264 Specification.
//  * SEI payload alignment bits as described in Annex D must be included at the
//  * end of the payload if needed.
//  * The payload should not be NAL-encapsulated.
//  * Payloads are written first in order of input, apart from in the case when HRD
//  * is enabled where payloads are written after the Buffering Period SEI. */

type X264SeiPayloadT struct {
	payload_size ffcommon.FInt
	payload_type ffcommon.FInt
	payload      *ffcommon.FUint8T
}

type X264SeiT struct {
	NumPayloads ffcommon.FInt
	Payloads    *X264SeiPayloadT
	/* In: optional callback to free each payload AND x264_sei_payload_t when used. */
	//void (*sei_free)( void* );
	SeiFree uintptr
}

type X264ImageT struct {
	ICsp    ffcommon.FInt        /* Colorspace */
	IPlane  ffcommon.FInt        /* Number of image planes */
	IStride [4]ffcommon.FInt     /* Strides for each plane */
	Plane   [4]*ffcommon.FUint8T /* Pointers to each plane */
}

type X264ImagePropertiesT struct {
	/* All arrays of data here are ordered as follows:
	 * each array contains one offset per macroblock, in raster scan order.  In interlaced
	 * mode, top-field MBs and bottom-field MBs are interleaved at the row level.
	 * Macroblocks are 16x16 blocks of pixels (with respect to the luma plane).  For the
	 * purposes of calculating the number of macroblocks, width and height are rounded up to
	 * the nearest 16.  If in interlaced mode, height is rounded up to the nearest 32 instead. */

	/* In: an array of quantizer offsets to be applied to this image during encoding.
	 *     These are added on top of the decisions made by x264.
	 *     Offsets can be fractional; they are added before QPs are rounded to integer.
	 *     Adaptive quantization must be enabled to use this feature.  Behavior if quant
	 *     offsets differ between encoding passes is undefined. */
	quant_offsets *ffcommon.FFloat
	/* In: optional callback to free quant_offsets when used.
	 *     Useful if one wants to use a different quant_offset array for each frame. */
	//void (*quant_offsets_free)( void* );
	quant_offsets_free uintptr

	/* In: optional array of flags for each macroblock.
	 *     Allows specifying additional information for the encoder such as which macroblocks
	 *     remain unchanged.  Usable flags are listed below.
	 *     x264_param_t.analyse.b_mb_info must be set to use this, since x264 needs to track
	 *     extra data internally to make full use of this information.
	 *
	 * Out: if b_mb_info_update is set, x264 will update this array as a result of encoding.
	 *
	 *      For "MBINFO_CONSTANT", it will remove this flag on any macroblock whose decoded
	 *      pixels have changed.  This can be useful for e.g. noting which areas of the
	 *      frame need to actually be blitted. Note: this intentionally ignores the effects
	 *      of deblocking for the current frame, which should be fine unless one needs exact
	 *      pixel-perfect accuracy.
	 *
	 *      Results for MBINFO_CONSTANT are currently only set for P-frames, and are not
	 *      guaranteed to enumerate all blocks which haven't changed.  (There may be false
	 *      negatives, but no false positives.)
	 */
	mb_info *ffcommon.FUint8T
	/* In: optional callback to free mb_info when used. */
	//void (*mb_info_free)( void* );
	mb_info_free uintptr

	/* The macroblock is constant and remains unchanged from the previous frame. */
	// const X264_MBINFO_CONSTANT   (1U<<0)
	/* More flags may be added in the future. */

	/* Out: SSIM of the the frame luma (if x264_param_t.b_ssim is set) */
	f_ssim ffcommon.FDouble
	/* Out: Average PSNR of the frame (if x264_param_t.b_psnr is set) */
	f_psnr_avg ffcommon.FDouble
	/* Out: PSNR of Y, U, and V (if x264_param_t.b_psnr is set) */
	f_psnr [3]ffcommon.FDouble

	/* Out: Average effective CRF of the encoded frame */
	f_crf_avg ffcommon.FDouble
}

const X264_MBINFO_CONSTANT = (1 << 0)

type X264PictureT struct {
	/* In: force picture type (if not auto)
	 *     If x264 encoding parameters are violated in the forcing of picture types,
	 *     x264 will correct the input picture type and log a warning.
	 * Out: type of the picture encoded */
	i_type ffcommon.FInt
	/* In: force quantizer for != X264_QP_AUTO */
	i_qpplus1 ffcommon.FInt
	/* In: pic_struct, for pulldown/doubling/etc...used only if b_pic_struct=1.
	 *     use pic_struct_e for pic_struct inputs
	 * Out: pic_struct element associated with frame */
	i_pic_struct ffcommon.FInt
	/* Out: whether this frame is a keyframe.  Important when using modes that result in
	 * SEI recovery points being used instead of IDR frames. */
	b_keyframe ffcommon.FInt
	/* In: user pts, Out: pts of encoded picture (user)*/
	IPts ffcommon.FInt64T
	/* Out: frame dts. When the pts of the first frame is close to zero,
	 *      initial frames may have a negative dts which must be dealt with by any muxer */
	i_dts ffcommon.FInt64T
	/* In: custom encoding parameters to be set from this frame forwards
	   (in coded order, not display order). If NULL, continue using
	   parameters from the previous frame.  Some parameters, such as
	   aspect ratio, can only be changed per-GOP due to the limitations
	   of H.264 itself; in this case, the caller must force an IDR frame
	   if it needs the changed parameter to apply immediately. */
	param *X264ParamT
	/* In: raw image data */
	/* Out: reconstructed image data.  x264 may skip part of the reconstruction process,
	   e.g. deblocking, in frames where it isn't necessary.  To force complete
	   reconstruction, at a small speed cost, set b_full_recon. */
	Img X264ImageT
	/* In: optional information to modify encoder decisions for this frame
	 * Out: information about the encoded frame */
	prop X264ImagePropertiesT
	/* Out: HRD timing information. Output only when i_nal_hrd is set. */
	hrd_timing X264HrdT
	/* In: arbitrary user SEI (e.g subtitles, AFDs) */
	extra_sei X264SeiT
	/* private user data. copied from input to output frames. */
	opaque ffcommon.FVoidP
}

// /* x264_picture_init:
//   - initialize an x264_picture_t.  Needs to be done if the calling application
//   - allocates its own x264_picture_t as opposed to using x264_picture_alloc. */
//
// X264_API void x264_picture_init( x264_picture_t *pic );
func (pic *X264PictureT) X264PictureInit() {
	libx264common.GetLibx264Dll().NewProc("x264_picture_init").Call(
		uintptr(unsafe.Pointer(pic)),
	)
}

// /* x264_picture_alloc:
//   - alloc data for a picture. You must call x264_picture_clean on it.
//   - returns 0 on success, or -1 on malloc failure or invalid colorspace. */
//
// X264_API int x264_picture_alloc( x264_picture_t *pic, int i_csp, int i_width, int i_height );
func (pic *X264PictureT) X264PictureAlloc(i_csp, i_width, i_height ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := libx264common.GetLibx264Dll().NewProc("x264_picture_alloc").Call(
		uintptr(unsafe.Pointer(pic)),
		uintptr(i_csp),
		uintptr(i_width),
		uintptr(i_height),
	)
	res = ffcommon.FInt(t)
	return
}

// /* x264_picture_clean:
//   - free associated resource for a x264_picture_t allocated with
//   - x264_picture_alloc ONLY */
//
// X264_API void x264_picture_clean( x264_picture_t *pic );
func (pic *X264PictureT) X264PictureClean() {
	libx264common.GetLibx264Dll().NewProc("x264_picture_clean").Call(
		uintptr(unsafe.Pointer(pic)),
	)
}

// /****************************************************************************
//  * Encoder functions
//  ****************************************************************************/

/* Force a link error in the case of linking against an incompatible API version.
 * Glue consts exist to force correct macro expansion; the final output of the macro
 * is x264_encoder_open_##X264_BUILD (for purposes of dlopen). */
// const x264_encoder_glue1(x,y) x##y
// const x264_encoder_glue2(x,y) x264_encoder_glue1(x,y)
// const x264_encoder_open x264_encoder_glue2(x264_encoder_open_,X264_BUILD)

// /* x264_encoder_open:
//   - create a new encoder handler, all parameters from x264_param_t are copied */
//
// X264_API x264_t *x264_encoder_open( x264_param_t * );
func (this *X264ParamT) X264EncoderOpen164() (res *X264T) {
	t, _, _ := libx264common.GetLibx264Dll().NewProc("x264_encoder_open_164").Call(
		uintptr(unsafe.Pointer(this)),
	)
	res = (*X264T)(unsafe.Pointer(t))
	return
}

// /* x264_encoder_reconfig:
//   - various parameters from x264_param_t are copied.
//   - this takes effect immediately, on whichever frame is encoded next;
//   - due to delay, this may not be the next frame passed to encoder_encode.
//   - if the change should apply to some particular frame, use x264_picture_t->param instead.
//   - returns 0 on success, negative on parameter validation error.
//   - not all parameters can be changed; see the actual function for a detailed breakdown.
//     *
//   - since not all parameters can be changed, moving from preset to preset may not always
//   - fully copy all relevant parameters, but should still work usably in practice. however,
//   - more so than for other presets, many of the speed shortcuts used in ultrafast cannot be
//   - switched out of; using reconfig to switch between ultrafast and other presets is not
//   - recommended without a more fine-grained breakdown of parameters to take this into account. */
//
// X264_API int x264_encoder_reconfig( x264_t *, x264_param_t * );
func (this *X264T) X264EncoderReconfig(p1 *X264ParamT) (res ffcommon.FInt) {
	t, _, _ := libx264common.GetLibx264Dll().NewProc("x264_encoder_reconfig").Call(
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(p1)),
	)
	res = ffcommon.FInt(t)
	return
}

// /* x264_encoder_parameters:
//   - copies the current internal set of parameters to the pointer provided
//   - by the caller.  useful when the calling application needs to know
//   - how x264_encoder_open has changed the parameters, or the current state
//   - of the encoder after multiple x264_encoder_reconfig calls.
//   - note that the data accessible through pointers in the returned param struct
//   - (e.g. filenames) should not be modified by the calling application. */
//
// X264_API void x264_encoder_parameters( x264_t *, x264_param_t * );
func (this *X264T) X264EncoderParameters(p1 *X264ParamT) {
	libx264common.GetLibx264Dll().NewProc("x264_encoder_parameters").Call(
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(p1)),
	)
}

// /* x264_encoder_headers:
//   - return the SPS and PPS that will be used for the whole stream.
//   - *pi_nal is the number of NAL units outputted in pp_nal.
//   - returns the number of bytes in the returned NALs.
//   - returns negative on error.
//   - the payloads of all output NALs are guaranteed to be sequential in memory. */
//
// X264_API int x264_encoder_headers( x264_t *, x264_nal_t **pp_nal, int *pi_nal );
func (this *X264T) X264EncoderHeaders(pp_nal **X264NalT, pi_nal *ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := libx264common.GetLibx264Dll().NewProc("x264_encoder_headers").Call(
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pp_nal)),
		uintptr(unsafe.Pointer(pi_nal)),
	)
	res = ffcommon.FInt(t)
	return
}

// /* x264_encoder_encode:
//   - encode one picture.
//   - *pi_nal is the number of NAL units outputted in pp_nal.
//   - returns the number of bytes in the returned NALs.
//   - returns negative on error and zero if no NAL units returned.
//   - the payloads of all output NALs are guaranteed to be sequential in memory. */
//
// X264_API int x264_encoder_encode( x264_t *, x264_nal_t **pp_nal, int *pi_nal, x264_picture_t *pic_in, x264_picture_t *pic_out );
func (this *X264T) X264EncoderEncode(pp_nal **X264NalT, pi_nal *ffcommon.FInt, pic_in, pic_out *X264PictureT) (res ffcommon.FInt) {
	t, _, _ := libx264common.GetLibx264Dll().NewProc("x264_encoder_encode").Call(
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(pp_nal)),
		uintptr(unsafe.Pointer(pi_nal)),
		uintptr(unsafe.Pointer(pic_in)),
		uintptr(unsafe.Pointer(pic_out)),
	)
	res = ffcommon.FInt(t)
	return
}

// /* x264_encoder_close:
//   - close an encoder handler */
//
// X264_API void x264_encoder_close( x264_t * );
func (this *X264T) X264EncoderClose() {
	libx264common.GetLibx264Dll().NewProc("x264_encoder_close").Call(
		uintptr(unsafe.Pointer(this)),
	)
}

// /* x264_encoder_delayed_frames:
//   - return the number of currently delayed (buffered) frames
//   - this should be used at the end of the stream, to know when you have all the encoded frames. */
//
// X264_API int x264_encoder_delayed_frames( x264_t * );
func (this *X264T) X264EncoderDelayedFrames() (res ffcommon.FInt) {
	t, _, _ := libx264common.GetLibx264Dll().NewProc("x264_encoder_delayed_frames").Call(
		uintptr(unsafe.Pointer(this)),
	)
	res = ffcommon.FInt(t)
	return
}

// /* x264_encoder_maximum_delayed_frames( x264_t * ):
//   - return the maximum number of delayed (buffered) frames that can occur with the current
//   - parameters. */
//
// X264_API int x264_encoder_maximum_delayed_frames( x264_t * );
func (this *X264T) X264EncoderMaximumDelayedFrames() (res ffcommon.FInt) {
	t, _, _ := libx264common.GetLibx264Dll().NewProc("x264_encoder_maximum_delayed_frames").Call(
		uintptr(unsafe.Pointer(this)),
	)
	res = ffcommon.FInt(t)
	return
}

// /* x264_encoder_intra_refresh:
//   - If an intra refresh is not in progress, begin one with the next P-frame.
//   - If an intra refresh is in progress, begin one as soon as the current one finishes.
//   - Requires that b_intra_refresh be set.
//     *
//   - Useful for interactive streaming where the client can tell the server that packet loss has
//   - occurred.  In this case, keyint can be set to an extremely high value so that intra refreshes
//   - only occur when calling x264_encoder_intra_refresh.
//     *
//   - In multi-pass encoding, if x264_encoder_intra_refresh is called differently in each pass,
//   - behavior is undefined.
//     *
//   - Should not be called during an x264_encoder_encode. */
//
// X264_API void x264_encoder_intra_refresh( x264_t * );
func (this *X264T) X264EncoderIntraRefresh() {
	libx264common.GetLibx264Dll().NewProc("x264_encoder_intra_refresh").Call(
		uintptr(unsafe.Pointer(this)),
	)
}

// /* x264_encoder_invalidate_reference:
//   - An interactive error resilience tool, designed for use in a low-latency one-encoder-few-clients
//   - system.  When the client has packet loss or otherwise incorrectly decodes a frame, the encoder
//   - can be told with this command to "forget" the frame and all frames that depend on it, referencing
//   - only frames that occurred before the loss.  This will force a keyframe if no frames are left to
//   - reference after the aforementioned "forgetting".
//     *
//   - It is strongly recommended to use a large i_dpb_size in this case, which allows the encoder to
//   - keep around extra, older frames to fall back on in case more recent frames are all invalidated.
//   - Unlike increasing i_frame_reference, this does not increase the number of frames used for motion
//   - estimation and thus has no speed impact.  It is also recommended to set a very large keyframe
//   - interval, so that keyframes are not used except as necessary for error recovery.
//     *
//   - x264_encoder_invalidate_reference is not currently compatible with the use of B-frames or intra
//   - refresh.
//     *
//   - In multi-pass encoding, if x264_encoder_invalidate_reference is called differently in each pass,
//   - behavior is undefined.
//     *
//   - Should not be called during an x264_encoder_encode, but multiple calls can be made simultaneously.
//     *
//   - Returns 0 on success, negative on failure. */
//
// X264_API int x264_encoder_invalidate_reference( x264_t *, int64_t pts );
func (this *X264T) X264EncoderInvalidateReference(pts ffcommon.FInt64T) (res ffcommon.FInt) {
	t, _, _ := libx264common.GetLibx264Dll().NewProc("x264_encoder_invalidate_reference").Call(
		uintptr(unsafe.Pointer(this)),
		uintptr(pts),
	)
	res = ffcommon.FInt(t)
	return
}

// #ifdef __cplusplus
// }
// #endif

// #endif
