package entities

import (
	"encoding/json"
	"strings"
)

type Weapon struct {
	UUID            string `json:"uuid"`
	DisplayName     string `json:"displayName"`
	Category        string `json:"category"`
	DefaultSkinUUID string `json:"defaultSkinUuid"`
	DisplayIcon     string `json:"displayIcon"`
	KillStreamIcon  string `json:"killStreamIcon"`
	AssetPath       string `json:"assetPath"`
	WeaponStats     struct {
		FireRate            float64 `json:"fireRate"`
		MagazineSize        int     `json:"magazineSize"`
		RunSpeedMultiplier  float64 `json:"runSpeedMultiplier"`
		EquipTimeSeconds    float64 `json:"equipTimeSeconds"`
		ReloadTimeSeconds   float64 `json:"reloadTimeSeconds"`
		FirstBulletAccuracy float64 `json:"firstBulletAccuracy"`
		ShotgunPelletCount  int     `json:"shotgunPelletCount"`
		WallPenetration     string  `json:"wallPenetration"`
		Feature             string  `json:"feature"`
		FireMode            string  `json:"fireMode"`
		AltFireType         string  `json:"altFireType"`
		AdsStats            struct {
			ZoomMultiplier      float64 `json:"zoomMultiplier"`
			FireRate            float64 `json:"fireRate"`
			RunSpeedMultiplier  float64 `json:"runSpeedMultiplier"`
			BurstCount          int     `json:"burstCount"`
			FirstBulletAccuracy float64 `json:"firstBulletAccuracy"`
		} `json:"adsStats"`
		AltShotgunStats struct {
			ShotgunPelletCount int     `json:"shotgunPelletCount"`
			BurstRate          float64 `json:"burstRate"`
		} `json:"altShotgunStats"`
		AirBurstStats struct {
			ShotgunPelletCount int     `json:"shotgunPelletCount"`
			BurstDistance      float64 `json:"burstDistance"`
		} `json:"airBurstStats"`
		DamageRanges []struct {
			RangeStartMeters int     `json:"rangeStartMeters"`
			RangeEndMeters   int     `json:"rangeEndMeters"`
			HeadDamage       float64 `json:"headDamage"`
			BodyDamage       float64 `json:"bodyDamage"`
			LegDamage        float64 `json:"legDamage"`
		} `json:"damageRanges"`
	} `json:"weaponStats"`
	ShopData struct {
		Cost         int    `json:"cost"`
		Category     string `json:"category"`
		CategoryText string `json:"categoryText"`
		GridPosition struct {
			Row    int `json:"row"`
			Column int `json:"column"`
		} `json:"gridPosition"`
		CanBeTrashed bool   `json:"canBeTrashed"`
		Image        string `json:"image"`
		NewImage     string `json:"newImage"`
		NewImage2    string `json:"newImage2"`
		AssetPath    string `json:"assetPath"`
	} `json:"shopData"`
	Skins []struct {
		UUID            string `json:"uuid"`
		DisplayName     string `json:"displayName"`
		ThemeUUID       string `json:"themeUuid"`
		ContentTierUUID string `json:"contentTierUuid"`
		DisplayIcon     string `json:"displayIcon"`
		Wallpaper       string `json:"wallpaper"`
		AssetPath       string `json:"assetPath"`
		Chromas         []struct {
			UUID          string `json:"uuid"`
			DisplayName   string `json:"displayName"`
			DisplayIcon   string `json:"displayIcon"`
			FullRender    string `json:"fullRender"`
			Swatch        string `json:"swatch"`
			StreamedVideo string `json:"streamedVideo"`
			AssetPath     string `json:"assetPath"`
		} `json:"chromas"`
		Levels []struct {
			UUID          string `json:"uuid"`
			DisplayName   string `json:"displayName"`
			LevelItem     string `json:"levelItem"`
			DisplayIcon   string `json:"displayIcon"`
			StreamedVideo string `json:"streamedVideo"`
			AssetPath     string `json:"assetPath"`
		} `json:"levels"`
	} `json:"skins"`
}

func (w *Weapon) UnmarshalJSON(data []byte) error {
	type Alias Weapon
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(w),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	w.UUID = strings.ToUpper(w.UUID)
	return nil
}
