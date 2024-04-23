package models

import "time"

type CS struct {
	DI_KEY        int
	DI_DT         int
	DI_SUBS_DI    *int64
	DI_REVISION   *int64
	DI_ACTIVE     *int64
	DI_EDIT_TIME  *int64
	DI_FLAG       *int64
	DI_REF        string
	DI_DATE       time.Time
	DI_CRE_DATE   time.Time
	DI_CRE_BY     string
	DI_CRE_CPTN   string
	DI_CRE_LGNN   string
	DI_UPD_DATE   time.Time
	DI_UPD_BY     string
	DI_UPD_CPTN   string
	DI_UPD_LGNN   string
	DI_DEL_DATE   *time.Time
	DI_DEL_BY     *string
	DI_DEL_CPTN   *string
	DI_DEL_LGNN   *string
	DI_PRN_TIME   int16
	DI_PRN_DATE   *time.Time
	DI_PRN_BY     *string
	DI_PRN_CPTN   *string
	DI_PRN_LGNN   *string
	DI_PRN_STATUS *int64
	DI_EXM_DATE   *time.Time
	DI_EXM_BY     *string
	DI_EXM_CPTN   *string
	DI_EXM_LGNN   *string
	DI_APV_DATE   *time.Time
	DI_APV_BY     *string
	DI_APV_CPTN   *string
	DI_APV_LGNN   *string
	DI_APV_STATUS *int64
	DI_DFRS       *int64
	DI_1ST_ITEMS  *int64
	DI_1ST_AMOUNT float64 // Assuming you handle money as float64
	DI_ITEMS      *int64
	DI_AMOUNT     float64 // Assuming you handle money as float64
	DI_AUTO       *int64
	DI_CREATOR_DI *int64
	DI_REMARK     *string
	DI_LASTUPD    *string

	DT_KEY           int
    DT_DOCCODE       string
    DT_THAIDESC      string
    DT_ENGDESC       *string
    DT_PROPERTIES    *int64
    DT_RUNTYPE       int16
    DT_PREFIX        string
    DT_DIGIT         int16
    DT_BOOKSIZE      *int64
    DT_ACCESS        int16
    DT_CANEDIT       string
    DT_DUPTYPE       *int64
    DT_RPF_CODE      *string
    DT_SHOW_SINCE    *int64
    DT_SHOW_NEXT     *int64
    DT_NEW_SINCE     *int64
    DT_NEW_NEXT      *int64
    DT_ENABLE        string
    DT_PREVIEW       *int64
    DT_NEED_APV      *int64
    DT_SELECT_PRT    *int64
    DT_APV_POS       *int64
    DT_EDIT_AF_POS   *int64
    DT_EDIT_PRT_POS  *int64
    DT_CHANGE_POS    *int64
    DT_CANCEL_POS    *int64
    DT_REPRT_POS     *int64
    DT_REF_B4_PRINT  string
    DT_PRINT_CANCEL  string
    DT_REPRT_MSG     *string
    DT_VAT_PREFIX    *string
    DT_LASTUPD       *string
    DT_BRANCH        *string

	BSTM_KEY          int
    BSTM_BNKAC        int
    BSTM_DI           *int64
    BSTM_TYPE         *int64
    BSTM_CHEQUE_NO    *string
    BSTM_CHEQUE_DD    *time.Time
    BSTM_RECNL_BSTMP  *int64
    BSTM_RECNL_DD     *time.Time
    BSTM_RECNL_SEQ    *int64
    BSTM_SHOW_ORDER   *int64
    BSTM_DEBIT        float64 // Assuming you handle money as float64
    BSTM_CREDIT       float64 // Assuming you handle money as float64
    BSTM_A_DEBIT      float64 // Assuming you handle money as float64
    BSTM_A_CREDIT     float64 // Assuming you handle money as float64
    BSTM_REMARK       *string
    BSTM_LASTUPD      *string

	CASHB_KEY           int
    CASHB_CASHAC        int
    CASHB_DI            *int64
    CASHB_TYPE          *int64
    CASHB_RECNL_CASHP   *int64
    CASHB_RECNL_SEQ     *int64
    CASHB_SHOW_ORDER    *int64
    CASHB_DEBIT         float64 // Assuming you handle money as float64
    CASHB_CREDIT        float64 // Assuming you handle money as float64
    CASHB_A_DEBIT       float64 // Assuming you handle money as float64
    CASHB_A_CREDIT      float64 // Assuming you handle money as float64
    CASHB_REMARK        *string
    CASHB_LASTUPD       *string

    TPH_KEY           int
    TPH_DI            int
    TPH_DEPT          int
    TPH_AR            int
    TPH_AP            int
    TPH_PSTATUS       *int // Use a pointer to int to handle NULL values
    TPH_LSTATUS       *int // Use a pointer to int to handle NULL values
    TPH_APPT_DATE     *time.Time // Use a pointer to time.Time to handle NULL values
    TPH_OPTION        *int // Use a pointer to int to handle NULL values
    TPH_REMARK        string
    TPH_REFER_XREF    string
    TPH_REFER_IREF    string
    TPH_REFER_PERSON  string
    TPH_REFER_XTRA1  string
    TPH_REFER_XTRA2  string
    TPH_TAP           *int // Use a pointer to int to handle NULL values
    TPH_BR            int
    TPH_PRJ           int
    TPH_MKTP          int
    TPH_PRMT          int
    TPH_SB            int
    TPH_SHIP_REMARK   string
    TPH_SHIP_DATE     time.Time
    TPH_SHIP_ADDB     *int // Use a pointer to int to handle NULL values
    TPH_CANCEL_DATE   *time.Time // Use a pointer to time.Time to handle NULL values
    TPH_SLMN          int
    TPH_WHT_TYPE      *int // Use a pointer to int to handle NULL values
    TPH_LASTUPD       string

    TPD_KEY           int
    TPD_TPH           int
    TPD_PMT           int
    TPD_SEQ           *int
    TPD_BAHT          float64 // Assuming you handle money as float64
    TPD_CARD_NO       *string
    TPD_EXPIRE        *string
    TPD_APPROV        *string
    TPD_CQIN_OWNER    *string
    TPD_CQIN_BANK     *int
    TPD_CQIN_BRANCH   *string
    TPD_CQIN_CHEQUE_NO *string
    TPD_CHEQUE_DD     *time.Time
    TPD_CHEQUE_AUTO   *int
    TPD_CQBK          *int
    TPD_CASHAC        *int
    TPD_BNKAC         *int
    TPD_AC            int
    TPD_REFER_REF     *string
    TPD_REFER_DATE    *time.Time
    TPD_REFER_DI      *int
    TPD_REFER_PPO     *int
    TPD_REFER_PPI     *int
    TPD_LASTUPD       *string

    ARP_KEY    int
    ARP_DI     int
    ARP_ARD    int
    ARP_PMT    int
    ARP_BAHT   float64 // Assuming you handle money as float64
    ARP_CRIN   *int
    ARP_CQIN   *int
    ARP_BNKAC  *int
    ARP_LASTUPD *string

    TRJ_KEY      int
    TRJ_DI       int
    TRJ_SEQ      *int
    TRJ_AC       int
    TRJ_DEPT     int
    TRJ_DEBIT    float64 // Assuming you handle money as float64
    TRJ_CREDIT   float64 // Assuming you handle money as float64
    TRJ_REMARK   *string
    TRJ_AUTO     string
    TRJ_IS_COST  string
    TRJ_LASTUPD  *string

    TRH_KEY          int
    TRH_DI           int
    TRH_PSTATUS      *int
    TRH_LSTATUS      *int
    TRH_FLAG         *int
    TRH_DEPT         int
    TRH_BR           int
    TRH_MKTP         int
    TRH_PRMT         int
    TRH_ARPRB        *string
    TRH_N_QTY        float64 // Assuming you handle money as float64
    TRH_N_ITEMS      *int
    TRH_REMARK       *string
    TRH_SB           int
    TRH_SHIP_REMARK  *string
    TRH_SHIP_DATE    time.Time
    TRH_SHIP_ADDB    *int
    TRH_VAT_TY       int
    TRH_VAT          float64 // Assuming you handle money as float64
    TRH_VAT_R        float64 // Assuming you handle money as float64
    TRH_VATIO        *int
    TRH_PRJ          int
    TRH_CANCEL_DATE  *time.Time
    TRH_LAST_DATE    *time.Time
    TRH_OPTION       *int
    TRH_REFER_XREF   *string
    TRH_REFER_IREF   *string
    TRH_REFER_PERSON *string
    TRH_REFER_XTRA1  *string
    TRH_REFER_XTRA2  *string
    TRH_TAP          *int
    TRH_LASTUPD      *string

    TRD_KEY           int
    TRD_TRH           int
    TRD_SEQ           *int
    TRD_GOODS         int
    TRD_KEYIN         string
    TRD_VAT_TY        int
    TRD_VAT           float64 // Assuming you handle money as float64
    TRD_VAT_R         float64 // Assuming you handle money as float64
    TRD_NM_PRC        float64 // Assuming you handle money as float64
    TRD_NM_VATIO      *int
    TRD_NM_U_DSC      *string
    TRD_QTY           float64 // Assuming you handle money as float64
    TRD_Q_FREE        float64 // Assuming you handle money as float64
    TRD_K_U_PRC       float64 // Assuming you handle money as float64
    TRD_U_PRC         float64 // Assuming you handle money as float64
    TRD_U_VATIO       *int
    TRD_DSC_KEYIN     *string
    TRD_DSC_KEYINV    float64 // Assuming you handle money as float64
    TRD_G_KEYIN       float64 // Assuming you handle money as float64
    TRD_G_SELL        float64 // Assuming you handle money as float64
    TRD_G_VAT         float64 // Assuming you handle money as float64
    TRD_G_AMT         float64 // Assuming you handle money as float64
    TRD_TDSC_KEYINV   float64 // Assuming you handle money as float64
    TRD_N_SELL        float64 // Assuming you handle money as float64
    TRD_N_VAT         float64 // Assuming you handle money as float64
    TRD_N_AMT         float64 // Assuming you handle money as float64
    TRD_B_SELL        float64 // Assuming you handle money as float64
    TRD_B_VAT         float64 // Assuming you handle money as float64
    TRD_B_AMT         float64 // Assuming you handle money as float64
    TRD_B_GROUP       *int
    TRD_B_RELOCATE    float64 // Assuming you handle money as float64
    TRD_B_UPRC        float64 // Assuming you handle money as float64
    TRD_SKU           int
    TRD_UTQQTY        float64 // Assuming you handle money as float64
    TRD_UTQNAME       string
    TRD_WEIGHT        float64 // Assuming you handle money as float64
    TRD_COST_TY       int
    TRD_WH_TY         int
    TRD_WH_RATE       float64 // Assuming you handle money as float64
    TRD_WH_TAX        float64 // Assuming you handle money as float64
    TRD_LOT_NO        *string
    TRD_SERIAL        *string
    TRD_EXP_D         *time.Time
    TRD_MAN_D         *time.Time
    TRD_WL            *int
    TRD_TO_WL         *int
    TRD_SH_AUTO       string
    TRD_SH_CODE       *string
    TRD_SH_NAME       *string
    TRD_SH_QTY        float64 // Assuming you handle money as float64
    TRD_SH_UPRC       float64 // Assuming you handle money as float64
    TRD_SH_GSELL      float64 // Assuming you handle money as float64
    TRD_SH_GVAT       float64 // Assuming you handle money as float64
    TRD_SH_GAMT       float64 // Assuming you handle money as float64
    TRD_SH_REMARK     *string
    TRD_RTN_UPRC      float64 // Assuming you handle money as float64
    TRD_RTN_AMT       float64 // Assuming you handle money as float64
    TRD_COMM_RATE     float64 // Assuming you handle money as float64
    TRD_COMM_AMT      float64 // Assuming you handle money as float64
    TRD_AC            int
    TRD_UQTY          float64 // Assuming you handle money as float64
    TRD_UQ_FREE       float64 // Assuming you handle money as float64
    TRD_OPTION        *int
    TRD_REFER_REF     *string
    TRD_REFER_DATE    *time.Time
    TRD_REFER_DI      *int
    TRD_REFER_TRH     *int
    TRD_REFER_TRD     *int
    TRD_OR_QTY        float64 // Assuming you handle money as float64
    TRD_OR_U_PRC      float64 // Assuming you handle money as float64
    TRD_OR_DSC_KEYINV float64 // Assuming you handle money as float64
    TRD_OR_G_KEYIN    float64 // Assuming you handle money as float64
    TRD_OR_Q_FREE     float64 // Assuming you handle money as float64
    TRD_AC_QTY        float64 // Assuming you handle money as float64
    TRD_AC_G_KEYIN    float64 // Assuming you handle money as float64
    TRD_AC_Q_FREE     float64 // Assuming you handle money as float64
    TRD_NX_QTY        float64 // Assuming you handle money as float64
    TRD_NX_DSC_KEYIN  *string
    TRD_NX_DSC_KEYINV float64 // Assuming you handle money as float64
    TRD_NX_G_KEYIN    float64 // Assuming you handle money as float64
    TRD_NX_Q_FREE     float64 // Assuming you handle money as float64
    TRD_ADJ_PAMT      float64 // Assuming you handle money as float64
    TRD_ADJ_MAMT      float64 // Assuming you handle money as float64
    TRD_B_PCNT        float64 // Assuming you handle money as float64
    TRD_PRICE         *int
    TRD_EQ_FACTOR     *string
    TRD_EQ_V          float64 // Assuming you handle money as float64
    TRD_CAMPAIGN      *int
    TRD_ARCPGN        *int
    TRD_ARCPGN_C      *string
    TRD_C_U_PRC       float64 // Assuming you handle money as float64
    TRD_C_DSC         *string
    TRD_C_DSCV        float64 // Assuming you handle money as float64
    TRD_CONSIGN       *string
    TRD_CNSN_B4       string
    TRD_CREATOR       *int
    TRD_SKUALT        *int
    TRD_ASRTH         *int
    TRD_ICCOLOR       *int
    TRD_ASRT_QTY      float64 // Assuming you handle money as float64
    TRD_ASRT_FREE     float64 // Assuming you handle money as float64
    TRD_ASRT_U_PRC    float64 // Assuming you handle money as float64
    TRD_ASRT_U_DSC    *string
    TRD_LASTUPD       *string

    SKM_KEY        int
    SKM_SKU        int
    SKM_DI         int
    SKM_WL         int
    SKM_REFER_DI   *int
    SKM_LOT_NO     *string
    SKM_SERIAL     *string
    SKM_QTY        float64 // Assuming you handle money as float64
    SKM_COST       float64 // Assuming you handle money as float64
    SKM_SELL       float64 // Assuming you handle money as float64
    SKM_VAT        float64 // Assuming you handle money as float64
    SKM_B4_DISC    float64 // Assuming you handle money as float64
    SKM_AF_GP      float64 // Assuming you handle money as float64
    SKM_Q_FREE     float64 // Assuming you handle money as float64
    SKM_EXP_D      *time.Time
    SKM_MAN_D      *time.Time
    SKM_SHOW_ORDER *int
    SKM_AC         int
    SKM_LASTUPD    *string

    SLD_KEY        int
    SLD_SLMN       int
    SLD_DI         *int
    SLD_REFER_DI   *int
    SLD_SHOW_ORDER *int
    SLD_SELL       float64 // Assuming you handle money as float64
    SLD_COMMISSION float64 // Assuming you handle money as float64
    SLD_A_SELL     float64 // Assuming you handle money as float64
    SLD_A_COMMISSION float64 // Assuming you handle money as float64
    SLD_REMARK     *string
    SLD_LASTUPD    *string

    VAT_KEY         int
	VAT_DI          int
	VAT_TYPE        *int
	VAT_RATE        float64 // Assuming you handle money as float64
	VAT_PSTATUS     *int
	VAT_REF         string
	VAT_DATE        time.Time
	VAT_RFR_REF     *string
	VAT_RFR_DATE    time.Time
	VAT_SBM_OPT     *int
	VAT_SBM_VTP     *int
	VAT_SBM_SEQ     *int
	VAT_SBM_REF     *string
	VAT_DESC        *string
	VAT_SV          float64 // Assuming you handle money as float64
	VAT_VAT         float64 // Assuming you handle money as float64
	VAT_SNV         float64 // Assuming you handle money as float64
	VAT_A_SV        float64 // Assuming you handle money as float64
	VAT_A_VAT       float64 // Assuming you handle money as float64
	VAT_A_SNV       float64 // Assuming you handle money as float64
	VAT_REMARK      *string
	VAT_SHV         string
	VAT_SHV_PCNT    float64 // Assuming you handle money as float64
	VAT_SHV_CS      float64 // Assuming you handle money as float64
	VAT_SHV_CV      float64 // Assuming you handle money as float64
	VAT_SHV_NCS     float64 // Assuming you handle money as float64
	VAT_SHV_NCV     float64 // Assuming you handle money as float64
	VAT_SHV_A_CS    float64 // Assuming you handle money as float64
	VAT_SHV_A_CV    float64 // Assuming you handle money as float64
	VAT_SHV_A_NCS   float64 // Assuming you handle money as float64
	VAT_SHV_A_NCV   float64 // Assuming you handle money as float64
	VAT_KV_ENABLED  string
	VAT_KS_ENABLED  string
	VAT_K_VAT       float64 // Assuming you handle money as float64
	VAT_K_SV        float64 // Assuming you handle money as float64
	VAT_K_SNV       float64 // Assuming you handle money as float64
	VAT_B4D_CS      float64 // Assuming you handle money as float64
	VAT_B4D_CV      float64 // Assuming you handle money as float64
	VAT_B4D_NCS     float64 // Assuming you handle money as float64
	VAT_B4D_NCV     float64 // Assuming you handle money as float64
	VAT_CREATOR_VAT *int
	VAT_LASTUPD     *string

    ARD_KEY        int
	ARD_AR         int
	ARD_DI         *int
	ARD_REFER_DI   *int
	ARD_SHOW_ORDER *int
	ARD_ARCD       int
	ARD_G_SNV      float64 // Assuming you handle money as float64
	ARD_G_SV       float64 // Assuming you handle money as float64
	ARD_G_VAT      float64 // Assuming you handle money as float64
	ARD_G_KEYIN    float64 // Assuming you handle money as float64
	ARD_TDSC_KEYIN *string
	ARD_TDSC_KEYINV float64 // Assuming you handle money as float64
	ARD_N_SNV       float64 // Assuming you handle money as float64
	ARD_N_SV        float64 // Assuming you handle money as float64
	ARD_N_VAT       float64 // Assuming you handle money as float64
	ARD_N_AMT       float64 // Assuming you handle money as float64
	ARD_CRNCYCODE   *string
	ARD_XCHG        *string
	ARD_B_SV        float64 // Assuming you handle money as float64
	ARD_B_SNV       float64 // Assuming you handle money as float64
	ARD_B_VAT       float64 // Assuming you handle money as float64
	ARD_B_AMT       float64 // Assuming you handle money as float64
	ARD_A_SV        float64 // Assuming you handle money as float64
	ARD_A_SNV       float64 // Assuming you handle money as float64
	ARD_A_VAT       float64 // Assuming you handle money as float64
	ARD_A_AMT       float64 // Assuming you handle money as float64
	ARD_CN_RAMT     float64 // Assuming you handle money as float64
	ARD_CASH_DC     *string
	ARD_CASH_B4     time.Time
	ARD_CASH_V      float64 // Assuming you handle money as float64
	ARD_WH_TAX      float64 // Assuming you handle money as float64
	ARD_BIL_DA      time.Time
	ARD_DUE_DA      time.Time
	ARD_CHQ_DA      time.Time
	ARD_REMARK      *string
	ARD_BILL_ADDB   *int
	ARD_P_AMT       float64 // Assuming you handle money as float64
	ARD_Q_AMT       float64 // Assuming you handle money as float64
	ARD_BILL_DI     *int
	ARD_LASTPM_DI   *int
	ARD_DPS_1_DI    *int
	ARD_DPS_2_DI    *int
	ARD_DPS_1_A     float64 // Assuming you handle money as float64
	ARD_DPS_2_A     float64 // Assuming you handle money as float64
	ARD_DPS_A       float64 // Assuming you handle money as float64
	ARD_LASTUPD     *string
}
